package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

type PersonnelBasicInformation struct {
	ID            int        `gorm:"column:ID;primary_key;AUTO_INCREMENT" json:"ID"`       // 序号
	JobNumber     string     `gorm:"column:job_number;size:10;not null" json:"job_number"` // 工号
	Name          string     `gorm:"column:name;size:8" json:"name"`                       // 姓名
	Gender        string     `gorm:"column:gender;size:2" json:"gender"`                   // 性别
	Education     string     `gorm:"column:education;size:8" json:"education"`             // 学历
	MaritalStatus string     `gorm:"column:marital_status;size:4" json:"marital_status"`   // 婚姻状况
	DateOfBirth   *time.Time `gorm:"column:date_of_birth" json:"date_of_birth"`            // 出生日期
	DutyRank      string     `gorm:"column:duty_rank;size:4" json:"duty_rank"`             // 职务职级
	Position      string     `gorm:"column:position;size:20" json:"position"`              // 职务
	Rank          string     `gorm:"column:rank;size:20;not null" json:"rank"`             // 职级
	Department    string     `gorm:"column:department;size:20" json:"department"`          // 部门
	IsItPresent   bool       `gorm:"column:is_it_present;not null" json:"is_it_present"`   // 是否在场
	EntryDate     *time.Time `gorm:"column:entry_date;not null" json:"entry_date"`         // 入职日期
}

func init() {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		log.Println(fmt.Sprintf("创建orm引擎失败, %s", err))
		return
	}
	defer db.Close()
	//创建人事基本资料表
	var personnel_basic_information PersonnelBasicInformation

	createErr := db.
		Set("gorm:table_options", "charset=utf8").
		AutoMigrate(personnel_basic_information).
		Error

	if createErr != nil {
		log.Println(fmt.Sprintf("创建人事基本资料表失败, %s", createErr))
	}
}

// TableName 将PersonnelBasicInformation映射为personnel_basic_information
func (table PersonnelBasicInformation) TableName() string {
	return "personnel_basic_information"
}

func NewPersonnelBasicInformation() *PersonnelBasicInformation {
	table := new(PersonnelBasicInformation)

	return table
}

// GetPersonnelBasicInformations: 获取所有人事基本资料记录
func GetPersonnelBasicInformations(qs map[string]interface{}, joins, fields []string, sortby string, offset int64, limit int64) ([]PersonnelBasicInformation, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	if sortby != "" {
		db = db.Order(sortby)
	}
	for _, join := range joins {
		db = db.Joins(join)
	}
	var records []PersonnelBasicInformation

	query := db.
		Select(fields).
		Offset(offset).
		Limit(limit).
		Where(qs).
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取所有人事基本资料记录失败，%s", query.Error)
	}
	return records, nil
}

// GetPersonnelBasicInformationByCondition: 根据给定条件获取人事基本资料记录
func GetPersonnelBasicInformationByCondition(qs interface{}, args ...interface{}) ([]PersonnelBasicInformation, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var records []PersonnelBasicInformation

	query := db.
		Where(qs, args...).
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取人事基本资料记录失败，%s", query.Error)
	}
	return records, nil
}

// GetPersonnelBasicInformationByPK: 根据主键获取人事基本资料记录
func GetPersonnelBasicInformationByPK(ID int) (*PersonnelBasicInformation, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var record PersonnelBasicInformation

	query := db.
		Where("ID=?", ID).
		Find(&record)

	if query.Error != nil {
		return nil, fmt.Errorf("获取人事基本资料记录失败，%s", query.Error)
	}
	return &record, nil
}

// Update: 更新人事基本资料记录
func (table *PersonnelBasicInformation) Update() error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	//开启事务
	tx := db.Begin()
	up := tx.Save(&table)

	if up.Error != nil {
		tx.Rollback()
		return fmt.Errorf("更新人事基本资料失败, %s", up.Error)
	}
	tx.Commit()

	return nil
}

// UpdateByCondition: 根据指定条件更新人事基本资料记录
func (table *PersonnelBasicInformation) UpdateByCondition(qs string, args ...interface{}) error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()
	//开启事务
	tx := db.Begin()

	up := tx.
		Model(&PersonnelBasicInformation{}).
		Where(qs, args...).
		Update(&table)

	if up.Error != nil {
		tx.Rollback()
		return fmt.Errorf("更新人事基本资料失败, %s", up.Error)
	}
	tx.Commit()

	return nil
}

// ImportPersonnelBasicInformation: 导入人事基础资料
func ImportPersonnelBasicInformation() (string, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return "", fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	err = NewPersonnelBasicInformation().DeleteAll()
	if err != nil {
		return "", err
	}
	//读取全场人员
	file, err := os.Open("./全场人员.csv")
	if err != nil {
		return "", err
	}
	defer file.Close()

	var personnels [][]string

	reader := csv.NewReader(file)
	for {
		personnel, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		if len(personnel) < 13 {
			return "", fmt.Errorf("全场人员表列数小于13")
		}
		personnels = append(personnels, personnel)
	}
	//读取在场人员
	file, err = os.Open("./在场人员.csv")
	if err != nil {
		return "", err
	}
	var presents []string

	reader = csv.NewReader(file)
	for {
		present, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		if len(present) < 3 {
			return "", fmt.Errorf("在场人员表列数小于3")
		}
		number := strings.TrimSpace(present[2]) //取第三列为工号
		presents = append(presents, number)
	}
	count := 0      //统计全场数量
	presentNum := 0 //统计在场数量

	for _, p := range personnels {
		birth, err := time.Parse("2006-01-02 15:04:05", p[6])
		if err != nil {
			return "", fmt.Errorf("全场人员表第7列必须是日期,%s", err)
		}
		date, err := time.Parse("2006-01-02 15:04:05", p[12])
		if err != nil {
			return "", fmt.Errorf("全场人员表第13列必须是日期,%s", err)
		}
		personnel := PersonnelBasicInformation{
			JobNumber:     p[1],
			Name:          p[2],
			Gender:        p[3],
			Education:     p[4],
			MaritalStatus: p[5],
			DateOfBirth:   &birth,
			DutyRank:      p[7],
			Position:      p[8],
			Rank:          p[9],
			Department:    p[10],
			EntryDate:     &date,
		}
		//更新在场状态
		for _, number := range presents {
			if number == personnel.JobNumber {
				presentNum++
				personnel.IsItPresent = true
			}
		}
		save := db.Save(&personnel)
		if save.Error != nil {
			return "", fmt.Errorf("导入%s失败, %s", personnel.JobNumber, save.Error)
		}
		count++
	}
	return fmt.Sprintf("add:%d,present:%d", count, presentNum), nil
}

// Insert: 新建人事基本资料记录
func (table *PersonnelBasicInformation) Insert() error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	save := db.Save(&table)
	if save.Error != nil {
		return fmt.Errorf("新建人事基本资料记录失败, %s", save.Error)
	}
	return nil
}

// Delete: 删除人事基本资料记录
func (table PersonnelBasicInformation) Delete() error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	del := db.Delete(&table)
	if del.Error != nil {
		return fmt.Errorf("删除人事基本资料记录失败, %s", del.Error)
	}
	return nil
}

// DeleteByCondition: 根据指定条件删除人事基本资料记录
func (table PersonnelBasicInformation) DeleteByCondition(qs string, args ...interface{}) error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	del := db.
		Where(qs, args...).
		Delete(&table)

	if del.Error != nil {
		return fmt.Errorf("删除人事基本资料记录失败, %s", del.Error)
	}
	return nil
}

// DeleteAll: 删除人事基本资料所有记录
func (table PersonnelBasicInformation) DeleteAll() error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	del := db.Exec("TRUNCATE TABLE " + table.TableName())
	if del.Error != nil {
		return fmt.Errorf("删除人事基本资料所有记录失败, %s", del.Error)
	}
	return nil
}
