package models

import (
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

type LotteryStaffTemporaryTable struct {
	ID            int    `gorm:"column:ID;primary_key;AUTO_INCREMENT" json:"ID"`       // 序号
	StageNumber   int    `gorm:"column:stage_number;size:11" json:"stage_number"`      // 阶段编号
	JobNumber     string `gorm:"column:job_number;size:10;not null" json:"job_number"` // 工号
	Department    string `gorm:"column:department;size:20" json:"department"`          // 部门
	Name          string `gorm:"column:name;size:8" json:"name"`                       // 姓名
	IsItRedundant string `gorm:"column:is_it_redundant;size:2" json:"is_it_redundant"` // 是否重复
}

func init() {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		log.Println(fmt.Sprintf("创建orm引擎失败, %s", err))
		return
	}
	defer db.Close()
	//创建抽奖人员临时表表
	var lottery_staff_temporary_table LotteryStaffTemporaryTable

	createErr := db.
		Set("gorm:table_options", "charset=utf8").
		AutoMigrate(lottery_staff_temporary_table).
		Error

	if createErr != nil {
		log.Println(fmt.Sprintf("创建抽奖人员临时表表失败, %s", createErr))
	}
}

// TableName 将LotteryStaffTemporaryTable映射为lottery_staff_temporary_table
func (table LotteryStaffTemporaryTable) TableName() string {
	return "lottery_staff_temporary_table"
}

func NewLotteryStaffTemporaryTable() *LotteryStaffTemporaryTable {
	table := new(LotteryStaffTemporaryTable)

	return table
}

// GetLotteryStaffTemporaryTables: 获取所有抽奖人员临时表记录
func GetLotteryStaffTemporaryTables(qs map[string]interface{}, joins, fields []string, sortby string, offset int64, limit int64) ([]LotteryStaffTemporaryTable, error) {
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
	var records []LotteryStaffTemporaryTable

	query := db.
		Select(fields).
		Offset(offset).
		Limit(limit).
		Where(qs).
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取所有抽奖人员临时表记录失败，%s", query.Error)
	}
	return records, nil
}

// GetLotteryStaffTemporaryTableByCondition: 根据给定条件获取抽奖人员临时表记录
func GetLotteryStaffTemporaryTableByCondition(qs interface{}, args ...interface{}) ([]LotteryStaffTemporaryTable, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var records []LotteryStaffTemporaryTable

	query := db.
		Where(qs, args...).
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取抽奖人员临时表记录失败，%s", query.Error)
	}
	return records, nil
}

// GetLotteryStaffTemporaryTableByPK: 根据主键获取抽奖人员临时表记录
func GetLotteryStaffTemporaryTableByPK(ID int) (*LotteryStaffTemporaryTable, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var record LotteryStaffTemporaryTable

	query := db.
		Where("ID=?", ID).
		Find(&record)

	if query.Error != nil {
		return nil, fmt.Errorf("获取抽奖人员临时表记录失败，%s", query.Error)
	}
	return &record, nil
}

// Update: 更新抽奖人员临时表记录
func (table *LotteryStaffTemporaryTable) Update() error {
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
		return fmt.Errorf("更新抽奖人员临时表失败, %s", up.Error)
	}
	tx.Commit()

	return nil
}

// UpdateByCondition: 根据指定条件更新抽奖人员临时表记录
func (table *LotteryStaffTemporaryTable) UpdateByCondition(qs string, args ...interface{}) error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()
	//开启事务
	tx := db.Begin()

	up := tx.
		Model(&LotteryStaffTemporaryTable{}).
		Where(qs, args...).
		Update(&table)

	if up.Error != nil {
		tx.Rollback()
		return fmt.Errorf("更新抽奖人员临时表失败, %s", up.Error)
	}
	tx.Commit()

	return nil
}

// Insert: 新建抽奖人员临时表记录
func (table *LotteryStaffTemporaryTable) Insert() error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	save := db.Save(&table)
	if save.Error != nil {
		return fmt.Errorf("新建抽奖人员临时表记录失败, %s", save.Error)
	}
	return nil
}

// Delete: 删除抽奖人员临时表记录
func (table LotteryStaffTemporaryTable) Delete() error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	del := db.Delete(&table)
	if del.Error != nil {
		return fmt.Errorf("删除抽奖人员临时表记录失败, %s", del.Error)
	}
	return nil
}

// DeleteByCondition: 根据指定条件删除抽奖人员临时表记录
func (table LotteryStaffTemporaryTable) DeleteByCondition(qs string, args ...interface{}) error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	del := db.
		Where(qs, args...).
		Delete(&table)

	if del.Error != nil {
		return fmt.Errorf("删除抽奖人员临时表记录失败, %s", del.Error)
	}
	return nil
}

// DeleteAll: 删除抽奖人员临时表所有记录
func (table LotteryStaffTemporaryTable) DeleteAll() error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	del := db.Exec("TRUNCATE TABLE " + table.TableName())
	if del.Error != nil {
		return fmt.Errorf("删除抽奖人员临时表所有记录失败, %s", del.Error)
	}
	return nil
}
