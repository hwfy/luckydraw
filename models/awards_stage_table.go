package models

import (
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

type AwardsStageTable struct {
	ID            int    `gorm:"column:ID;primary_key;AUTO_INCREMENT" json:"ID"`       // 序号
	Name          string `gorm:"column:name;size:8" json:"name"`                       // 名称
	Awards        string `gorm:"column:awards;size:10" json:"awards"`                  // 奖项
	Quantity      int    `gorm:"column:quantity;size:10" json:"quantity"`              // 数量
	IsItPresent   string `gorm:"column:is_it_present;size:2" json:"is_it_present"`     // 是否在场
	IsItRedundant string `gorm:"column:is_it_redundant;size:2" json:"is_it_redundant"` // 是否重复
	IsItDisplayed bool   `gorm:"column:is_it_displayed" json:"is_it_displayed"`        // 是否显示

	AwardsConditionsTableID      []*AwardsConditionsTable     `gorm:"ForeignKey:stage_number;AssociationForeignKey:ID" json:"awards_conditions_table_ID"`
	LotteryStaffTemporaryTableID []LotteryStaffTemporaryTable `gorm:"ForeignKey:stage_number;AssociationForeignKey:ID" json:"lottery_staff_temporary_table_ID"`
	LotteryStaffQuantityTableID  []LotteryStaffQuantityTable  `gorm:"ForeignKey:stage_number;AssociationForeignKey:ID" json:"lottery_staff_quantity_table_ID"`
	WinningFormID                []WinningForm                `gorm:"ForeignKey:stage_number;AssociationForeignKey:ID" json:"winning_form_ID"`
}

func init() {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		log.Println(fmt.Sprintf("创建orm引擎失败, %s", err))
		return
	}
	defer db.Close()

	var awards_stage_table AwardsStageTable

	createErr := db.
		//创建奖项阶段表表
		Set("gorm:table_options", "charset=utf8").
		AutoMigrate(awards_stage_table).
		//创建奖项条件表、索引、外键
		AutoMigrate(&AwardsConditionsTable{}).
		AddIndex("idx_awards_conditions_table_stage_number", "stage_number").
		AddForeignKey("stage_number", "awards_stage_table(ID)", "RESTRICT", "RESTRICT").
		//创建抽奖人员表、索引、外键
		AutoMigrate(&LotteryStaffTemporaryTable{}).
		AddIndex("idx_lottery_staff_temporary_table_stage_number", "stage_number").
		AddForeignKey("stage_number", "awards_stage_table(ID)", "RESTRICT", "RESTRICT").
		//创建抽奖人员表、索引、外键
		AutoMigrate(&LotteryStaffQuantityTable{}).
		AddIndex("idx_lottery_staff_quantity_table_stage_number", "stage_number").
		AddForeignKey("stage_number", "awards_stage_table(ID)", "RESTRICT", "RESTRICT").
		//创建中将人员表、索引、外键
		AutoMigrate(&WinningForm{}).
		AddIndex("idx_winning_form_stage_number", "stage_number").
		AddForeignKey("stage_number", "awards_stage_table(ID)", "RESTRICT", "RESTRICT").
		Error

	if createErr != nil {
		log.Println(fmt.Sprintf("创建奖项阶段表表失败, %s", createErr))
	}
}

// TableName 将AwardsStageTable映射为awards_stage_table
func (table AwardsStageTable) TableName() string {
	return "awards_stage_table"
}

func NewAwardsStageTable() *AwardsStageTable {
	table := new(AwardsStageTable)

	return table
}

// GetAwardsStageTables: 获取所有奖项阶段表记录
func GetAwardsStageTables(qs map[string]interface{}, joins, fields []string, sortby string, offset int64, limit int64) ([]AwardsStageTable, error) {
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
	var records []AwardsStageTable

	query := db.
		Preload("AwardsConditionsTableID").
		Preload("LotteryStaffTemporaryTableID").
		Preload("LotteryStaffQuantityTableID").
		Preload("WinningFormID").
		Select(fields).
		Offset(offset).
		Limit(limit).
		Where(qs).
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取所有奖项阶段表记录失败，%s", query.Error)
	}
	return records, nil
}

// GetAwardsStageTableByCondition: 根据给定条件获取奖项阶段表记录
func GetAwardsStageTableByCondition(qs interface{}, args ...interface{}) ([]AwardsStageTable, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var records []AwardsStageTable

	query := db.
		Preload("AwardsConditionsTableID").
		Preload("LotteryStaffTemporaryTableID").
		Preload("LotteryStaffQuantityTableID").
		Preload("WinningFormID").
		Where(qs, args...).
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取奖项阶段表记录失败，%s", query.Error)
	}
	return records, nil
}

// GetAwardsStageTableByPK: 根据主键获取奖项阶段表记录
func GetAwardsStageTableByPK(ID int) (*AwardsStageTable, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var record AwardsStageTable

	query := db.
		Preload("AwardsConditionsTableID").
		Preload("LotteryStaffTemporaryTableID").
		Preload("LotteryStaffQuantityTableID").
		Preload("WinningFormID").
		Where("ID=?", ID).
		Find(&record)

	if query.Error != nil {
		return nil, fmt.Errorf("获取奖项阶段表记录失败，%s", query.Error)
	}
	return &record, nil
}

// Update: 更新奖项阶段表记录
func (table *AwardsStageTable) Update() error {
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
		return fmt.Errorf("更新奖项阶段表失败, %s", up.Error)
	}
	tx.Commit()

	return nil
}

// UpdateByCondition: 根据指定条件更新奖项阶段表记录
func (table *AwardsStageTable) UpdateByCondition(qs string, args ...interface{}) error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()
	//开启事务
	tx := db.Begin()

	up := tx.
		Model(&AwardsStageTable{}).
		Where(qs, args...).
		Update(&table)

	if up.Error != nil {
		tx.Rollback()
		return fmt.Errorf("更新奖项阶段表失败, %s", up.Error)
	}
	tx.Commit()

	return nil
}

// Insert: 新建奖项阶段表记录
func (table *AwardsStageTable) Insert() error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	save := db.Save(&table)
	if save.Error != nil {
		return fmt.Errorf("新建奖项阶段表记录失败, %s", save.Error)
	}
	return nil
}

// Delete: 删除奖项阶段表记录
func (table AwardsStageTable) Delete() error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()
	//删除所有相关抽奖条件
	err = NewAwardsConditionsTable().DeleteByCondition("stage_number=?", table.ID)
	if err != nil {
		return err
	}
	//删除所有相关抽奖人员
	err = NewLotteryStaffTemporaryTable().DeleteByCondition("stage_number=?", table.ID)
	if err != nil {
		return err
	}
	err = NewLotteryStaffQuantityTable().DeleteByCondition("stage_number=?", table.ID)
	if err != nil {
		return err
	}
	//删除所有相关中奖人员
	err = NewWinningForm().DeleteByCondition("stage_number=?", table.ID)
	if err != nil {
		return err
	}
	del := db.Delete(&table)
	if del.Error != nil {
		return fmt.Errorf("删除奖项阶段表记录失败, %s", del.Error)
	}
	return nil
}

// DeleteByCondition: 根据指定条件删除奖项阶段表记录
func (table AwardsStageTable) DeleteByCondition(qs string, args ...interface{}) error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	del := db.
		Where(qs, args...).
		Delete(&table)

	if del.Error != nil {
		return fmt.Errorf("删除奖项阶段表记录失败, %s", del.Error)
	}
	return nil
}
