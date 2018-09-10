package models

import (
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

type AwardsConditionsTable struct {
	ID               int      `gorm:"column:ID;primary_key;AUTO_INCREMENT" json:"ID"`            // 序号
	StageNumber      int      `gorm:"column:stage_number;size:11" json:"stage_number"`           // 阶段编号
	ConditionName    string   `gorm:"column:condition_name;size:10" json:"condition_name"`       // 条件名
	ConditionFlag    string   `json:"condition_flag"`                                            // 条件字段名
	ConditionSymbol  string   `gorm:"column:condition_symbol;size:4" json:"condition_symbol"`    // 条件符
	ConditionValue   string   `gorm:"column:condition_value;size:20" json:"condition_value"`     // 条件值
	ConditionType    string   `gorm:"column:condition_type;size:20" json:"condition_type"`       // 条件值类型
	ConditionDefault string   `gorm:"column:condition_default;size:20" json:"condition_default"` // 条件默认值
	ConditionValues  []string `gorm:"-" json:"column_values"`                                    // 条件可选值
}

func init() {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		log.Println(fmt.Sprintf("创建orm引擎失败, %s", err))
		return
	}
	defer db.Close()
	//创建奖项条件表表
	var awards_conditions_table AwardsConditionsTable

	createErr := db.
		Set("gorm:table_options", "charset=utf8").
		AutoMigrate(awards_conditions_table).
		Error

	if createErr != nil {
		log.Println(fmt.Sprintf("创建奖项条件表表失败, %s", createErr))
	}
}

// TableName 将AwardsConditionsTable映射为awards_conditions_table
func (table AwardsConditionsTable) TableName() string {
	return "awards_conditions_table"
}

func NewAwardsConditionsTable() *AwardsConditionsTable {
	table := new(AwardsConditionsTable)

	return table
}

// GetAwardsConditionsTables: 获取所有奖项条件表记录
func GetAwardsConditionsTables(qs map[string]interface{}, joins, fields []string, sortby string, offset int64, limit int64) ([]AwardsConditionsTable, error) {
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
	var records []AwardsConditionsTable

	query := db.
		Select(fields).
		Offset(offset).
		Limit(limit).
		Where(qs).
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取所有奖项条件表记录失败，%s", query.Error)
	}
	return records, nil
}

// GetAwardsConditionsTableByCondition: 根据给定条件获取奖项条件表记录
func GetAwardsConditionsTableByCondition(qs interface{}, args ...interface{}) ([]*AwardsConditionsTable, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var records []*AwardsConditionsTable

	query := db.
		Where(qs, args...).
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取奖项条件表记录失败，%s", query.Error)
	}
	//添加条件可选值
	personnel := NewPersonnelBasicInformation()

	for _, record := range records {
		values, err := getColumnValues(db, personnel.TableName(), record.ConditionFlag)
		if err != nil {
			return nil, err
		}
		record.ConditionValues = values
	}
	return records, nil
}

// GetAwardsConditionsTableByPK: 根据主键获取奖项条件表记录
func GetAwardsConditionsTableByPK(ID int) (*AwardsConditionsTable, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var record AwardsConditionsTable

	query := db.
		Where("ID=?", ID).
		Find(&record)

	if query.Error != nil {
		return nil, fmt.Errorf("获取奖项条件表记录失败，%s", query.Error)
	}
	return &record, nil
}

// Update: 更新奖项条件表记录
func (table *AwardsConditionsTable) Update() error {
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
		return fmt.Errorf("更新奖项条件表失败, %s", up.Error)
	}
	tx.Commit()

	return nil
}

// UpdateByCondition: 根据指定条件更新奖项条件表记录
func (table *AwardsConditionsTable) UpdateByCondition(qs string, args ...interface{}) error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()
	//开启事务
	tx := db.Begin()

	up := tx.
		Model(&AwardsConditionsTable{}).
		Where(qs, args...).
		Update(&table)

	if up.Error != nil {
		tx.Rollback()
		return fmt.Errorf("更新奖项条件表失败, %s", up.Error)
	}
	tx.Commit()

	return nil
}

// Insert: 新建奖项条件表记录
func (table *AwardsConditionsTable) Insert() error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	save := db.Save(&table)
	if save.Error != nil {
		return fmt.Errorf("新建奖项条件表记录失败, %s", save.Error)
	}
	return nil
}

// Delete: 删除奖项条件表记录
func (table AwardsConditionsTable) Delete() error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	del := db.Delete(&table)
	if del.Error != nil {
		return fmt.Errorf("删除奖项条件表记录失败, %s", del.Error)
	}
	return nil
}

// DeleteByCondition: 根据指定条件删除奖项条件表记录
func (table AwardsConditionsTable) DeleteByCondition(qs string, args ...interface{}) error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	del := db.
		Where(qs, args...).
		Delete(&table)

	if del.Error != nil {
		return fmt.Errorf("删除奖项条件表记录失败, %s", del.Error)
	}
	return nil
}
