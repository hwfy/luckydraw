package models

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type WinningForm struct {
	ID            int        `gorm:"column:ID;primary_key;AUTO_INCREMENT" json:"ID"`       // 序号
	StageNumber   int        `gorm:"column:stage_number;size:11" json:"stage_number"`      // 阶段编号
	JobNumber     string     `gorm:"column:job_number;size:10" json:"job_number"`          // 工号
	Name          string     `gorm:"column:name;size:10" json:"name"`                      // 姓名
	Department    string     `gorm:"column:department;size:20" json:"department"`          // 部门
	IsItRedundant string     `gorm:"column:is_it_redundant;size:2" json:"is_it_redundant"` // 是否重复
	WinningTime   *time.Time `gorm:"column:winning_time" json:"winning_time"`              // 中奖时间

}

func init() {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		log.Println(fmt.Sprintf("创建orm引擎失败, %s", err))
		return
	}
	defer db.Close()
	//创建中奖表表
	var winning_form WinningForm

	createErr := db.
		Set("gorm:table_options", "charset=utf8").
		AutoMigrate(winning_form).
		Error

	if createErr != nil {
		log.Println(fmt.Sprintf("创建中奖表表失败, %s", createErr))
	}
}

// TableName 将WinningForm映射为winning_form
func (table WinningForm) TableName() string {
	return "winning_form"
}

func NewWinningForm() *WinningForm {
	table := new(WinningForm)

	return table
}

// GetWinningForms: 获取所有中奖表记录
func GetWinningForms(qs map[string]interface{}, joins, fields []string, sortby string, offset int64, limit int64) ([]WinningForm, error) {
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
	var records []WinningForm

	query := db.
		Select(fields).
		Offset(offset).
		Limit(limit).
		Where(qs).
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取所有中奖表记录失败，%s", query.Error)
	}
	return records, nil
}

// GetWinningPersonelJobNumbers: 获取所有中奖人员工号
func GetWinningPersonelJobNumbers() ([]string, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var records []WinningForm

	query := db.
		Select("job_number").
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取中奖表记录失败，%s", query.Error)
	}
	var numbers []string

	for _, record := range records {
		numbers = append(numbers, record.JobNumber)
	}
	return numbers, nil
}

// GetWinningPersonelJobNumber: 根据奖项阶段所有人员工号
func GetWinningPersonelJobNumber(stage int) ([]string, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var records []WinningForm

	query := db.
		Select("job_number").
		Where("stage_number=?", stage).
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取中奖表记录失败，%s", query.Error)
	}
	var numbers []string

	for _, record := range records {
		numbers = append(numbers, record.JobNumber)
	}
	return numbers, nil
}

// GetWinningFormByCondition: 根据给定条件获取中奖表记录
func GetWinningFormByCondition(qs interface{}, args ...interface{}) ([]WinningForm, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var records []WinningForm

	query := db.
		Where(qs, args...).
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取中奖表记录失败，%s", query.Error)
	}
	return records, nil
}

// GetWinningFormByPK: 根据主键获取中奖表记录
func GetWinningFormByPK(ID int) (*WinningForm, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var record WinningForm

	query := db.
		Where("ID=?", ID).
		Find(&record)

	if query.Error != nil {
		return nil, fmt.Errorf("获取中奖表记录失败，%s", query.Error)
	}
	return &record, nil
}

// Update: 更新中奖表记录
func (table *WinningForm) Update() error {
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
		return fmt.Errorf("更新中奖表失败, %s", up.Error)
	}
	tx.Commit()

	return nil
}

// UpdateByCondition: 根据指定条件更新中奖表记录
func (table *WinningForm) UpdateByCondition(qs string, args ...interface{}) error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()
	//开启事务
	tx := db.Begin()

	up := tx.
		Model(&WinningForm{}).
		Where(qs, args...).
		Update(&table)

	if up.Error != nil {
		tx.Rollback()
		return fmt.Errorf("更新中奖表失败, %s", up.Error)
	}
	tx.Commit()

	return nil
}

// Insert: 新建中奖表记录
func (table *WinningForm) Insert() (err error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	save := db.Save(&table)

	if save.Error != nil {
		err = fmt.Errorf("新建中奖表记录失败, %s", save.Error)
		return
	}
	//如果不允许重复删除待抽奖人员
	if table.IsItRedundant == "否" {
		staff := NewLotteryStaffTemporaryTable()
		err = staff.DeleteByCondition("job_number=?", table.JobNumber)
	}
	return
}

// Delete: 删除中奖表记录
func (table WinningForm) Delete() error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	del := db.Delete(&table)

	if del.Error != nil {
		return fmt.Errorf("删除中奖表记录失败, %s", del.Error)
	}
	return nil
}

// DeleteByCondition: 根据指定条件删除中奖表记录
func (table WinningForm) DeleteByCondition(qs string, args ...interface{}) error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	del := db.
		Where(qs, args...).
		Delete(&table)

	if del.Error != nil {
		return fmt.Errorf("删除中奖表记录失败, %s", del.Error)
	}
	return nil
}
