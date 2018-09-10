package models

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

type AwardsTable struct {
	ID               int     `gorm:"column:ID;primary_key;AUTO_INCREMENT" json:"ID"`             // 序号
	Name             string  `gorm:"column:name;size:8" json:"name"`                             // 名称
	Amount           float64 `gorm:"column:amount;size:20" json:"amount"`                        // 金额
	DisplayTheNumber int     `gorm:"column:display_the_number;size:4" json:"display_the_number"` // 显示数量
	IsItLottery      bool    `gorm:"column:is_it_lottery;size:4" json:"is_it_lottery"`           // 是否抽奖
	IsItOver         bool    `gorm:"column:is_it_over;size:4" json:"is_it_over"`                 // 是否完成

	AwardsStageTableName []*AwardsStageTable `gorm:"ForeignKey:awards;AssociationForeignKey:name" json:"awards_stage_table_name"`
}

func init() {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		log.Println(fmt.Sprintf("创建orm引擎失败, %s", err))
		return
	}
	defer db.Close()

	var awards_table AwardsTable

	createErr := db.
		Set("gorm:table_options", "charset=utf8").
		//创建奖项表表
		AutoMigrate(awards_table).
		AddIndex("idx_awards_table_name", "name").
		//创建子表、索引、外键
		AutoMigrate(&AwardsStageTable{}).
		AddIndex("idx_awards_stage_table_awards", "awards").
		AddForeignKey("awards", "awards_table(name)", "RESTRICT", "RESTRICT").
		Error

	if createErr != nil {
		log.Println(fmt.Sprintf("创建奖项表表失败, %s", createErr))
	}
}

// TableName 将AwardsTable映射为awards_table
func (table AwardsTable) TableName() string {
	return "awards_table"
}

func NewAwardsTable() *AwardsTable {
	table := new(AwardsTable)

	return table
}

// GetAwardsTables: 获取所有奖项表记录
func GetAwardsTables(qs map[string]interface{}, joins, fields []string, sortby string, offset int64, limit int64) ([]AwardsTable, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	for _, join := range joins {
		db = db.Joins(join)
	}
	if sortby != "" {
		db = db.Order(sortby)
	}
	var records []AwardsTable

	query := db.
		Preload("AwardsStageTableName").
		Select(fields).
		Offset(offset).
		Limit(limit).
		Where(qs).
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取所有奖项表记录失败，%s", query.Error)
	}
	//手动添加二级子表数据
	err = setAwardsConditions(records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// setAwardsConditions 设置奖项相关数据
func setAwardsConditions(records []AwardsTable) error {
	for _, record := range records {
		for _, stage := range record.AwardsStageTableName {
			conditions, err := GetAwardsConditionsTableByCondition("stage_number=?", stage.ID)
			if err != nil {
				return err
			}
			stage.AwardsConditionsTableID = conditions

			staffs, err := GetLotteryStaffTemporaryTableByCondition("stage_number=?", stage.ID)
			if err != nil {
				return err
			}
			stage.LotteryStaffTemporaryTableID = staffs

			winnings, err := GetWinningFormByCondition("stage_number=?", stage.ID)
			if err != nil {
				return err
			}
			stage.WinningFormID = winnings
		}
	}
	return nil
}

// GetAwardsTableByCondition: 根据给定条件获取奖项表记录
func GetAwardsTableByCondition(qs interface{}, args ...interface{}) ([]AwardsTable, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var records []AwardsTable

	query := db.
		Preload("AwardsStageTableName").
		Where(qs, args...).
		Find(&records)

	if query.Error != nil {
		return nil, fmt.Errorf("获取奖项表记录失败，%s", query.Error)
	}
	return records, nil
}

// GetAwardsTableByName: 根据奖项名称获取抽奖人员,抽奖页面点[开始]触发
func GetAwardsTableByName(name string) ([]LotteryStaffQuantityTable, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var record AwardsTable

	query := db.
		Preload("AwardsStageTableName").
		Where("name=?", name).
		Find(&record)

	if query.Error != nil {
		return nil, fmt.Errorf("获取奖项表记录失败,%s", query.Error)
	}
	return getCurrentLotteryStaff(&record)
}

// getCurrentLotteryStaff 获取当前抽奖人员
func getCurrentLotteryStaff(record *AwardsTable) ([]LotteryStaffQuantityTable, error) {
	var staffs []LotteryStaffQuantityTable

	for _, stage := range record.AwardsStageTableName {
		stageStaffs, err := GetLotteryStaffQuantityTableByCondition("stage_number=?", stage.ID)
		if err != nil {
			return nil, err
		}
		staffs = append(staffs, stageStaffs...)
	}
	return staffs, nil
}

// GetAwardsName: 获取当前抽奖名称
func GetAwardsName() (string, error) {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return "", fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var record AwardsTable

	query := db.
		Where("is_it_lottery=true").
		Find(&record)

	if query.Error != nil {
		return "", fmt.Errorf("当前未存在待抽奖项,%s", query.Error)
	}
	//重置抽奖状态
	exc := db.Exec("UPDATE awards_table SET is_it_lottery=false")
	if exc.Error != nil {
		return "", fmt.Errorf("更新奖项表所有状态失败, %s", exc.Error)
	}
	return record.Name, nil
}

// Update: 更新奖项表记录,奖项选择中点[开始]触发
func (table *AwardsTable) Update() error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()
	//重置抽奖状态,updates不生效
	exc := db.Exec("UPDATE awards_table SET is_it_lottery=false")
	if exc.Error != nil {
		return fmt.Errorf("更新奖项表所有状态失败, %s", exc.Error)
	}
	//清空指定抽奖人员表
	err = NewLotteryStaffQuantityTable().DeleteAll()
	if err != nil {
		return err
	}
	//设置指定数量抽奖人员
	err = setSpecifyLotteryStaffs(table)
	if err != nil {
		return err
	}
	//更新奖项状态为待抽奖
	tx := db.Begin()

	up := tx.Save(&table)
	if up.Error != nil {
		tx.Rollback()
		return fmt.Errorf("更新奖项表状态失败, %s", up.Error)
	}
	tx.Commit()

	return nil
}

// setSpecifyLotteryStaffs 当前奖项阶段所有抽奖人员
func setSpecifyLotteryStaffs(record *AwardsTable) (err error) {
	for _, stage := range record.AwardsStageTableName {
		stageStaffs, err := GetLotteryStaffTemporaryTableByCondition("stage_number=?", stage.ID)
		if err != nil {
			return err
		}
		winnings, err := GetWinningPersonelJobNumber(stage.ID)
		if err != nil {
			return err
		}
		quantity := stage.Quantity - len(winnings) //当前抽奖阶段需要抽取的人数

		specifys, err := getSpecifyLotteryStaffs(stageStaffs, quantity, stage.IsItRedundant)
		if err != nil {
			return err
		}
		if quantity > len(specifys) {
			return fmt.Errorf("%s中 %s 可抽奖数量剩余%d个,请调整抽奖数量或者更改条件!", record.Name, stage.Name, len(specifys))
		}
		stage.LotteryStaffQuantityTableID = specifys
	}
	return nil
}

// UpdateByCondition: 根据指定条件更新奖项表记录
func (table *AwardsTable) UpdateByCondition(qs string, args ...interface{}) error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()
	//开启事务
	tx := db.Begin()

	up := tx.
		Model(&AwardsTable{}).
		Where(qs, args...).
		Update(&table)

	if up.Error != nil {
		tx.Rollback()
		return fmt.Errorf("更新奖项表失败, %s", up.Error)
	}
	tx.Commit()

	return nil
}

// InsertAwards: 新建奖项表记录,奖项保存触发
func InsertAwards(tables []*AwardsTable) error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()
	//清空抽奖人员表
	err = NewLotteryStaffTemporaryTable().DeleteAll()
	if err != nil {
		return err
	}
	tx := db.Begin()

	for _, table := range tables {
		err = table.setLotteryStaffs()
		if err != nil {
			tx.Rollback()
			return err
		}
		save := tx.Save(&table)
		if save.Error != nil {
			tx.Rollback()
			return fmt.Errorf("新建奖项表记录失败, %s", save.Error)
		}
	}
	tx.Commit()

	return nil
}

// setLotteryStaffs 设置当前奖项所有抽奖人员
func (award *AwardsTable) setLotteryStaffs() error {
	for _, awardsStage := range award.AwardsStageTableName {
		winnings, err := GetWinningPersonelJobNumber(awardsStage.ID)
		if err != nil {
			return err
		}
		if awardsStage.Quantity < len(winnings) {
			return fmt.Errorf("%s中 %s 设置的奖项数量不能小于%s已经中奖数量!", award.Name, awardsStage.Name, awardsStage.Name)
		}
		where, err := getAwardsStageWhere(awardsStage)
		if err != nil {
			return err
		}
		personnels, err := GetPersonnelBasicInformationByCondition(where)
		if err != nil {
			return err
		}
		if awardsStage.Quantity > len(personnels)+len(winnings) {
			return fmt.Errorf("%s中 %s 设置的奖项数量不能大于%s剩余抽奖数量!", award.Name, awardsStage.Name, awardsStage.Name)
		}
		awardsStage.LotteryStaffTemporaryTableID = getLotteryStaffsFromPersonnels(personnels, awardsStage.IsItRedundant)
		awardsStage.WinningFormID = nil
	}
	return nil
}

// getAwardsStageWhere 得到当前阶段抽奖条件
func getAwardsStageWhere(awardsStage *AwardsStageTable) (string, error) {
	var wheres []string

	if awardsStage.IsItPresent == "是" {
		wheres = append(wheres, "is_it_present=true")
	}
	if awardsStage.IsItRedundant == "否" {
		numbers, err := GetWinningPersonelJobNumbers()
		if err != nil {
			return "", err
		}
		if len(numbers) != 0 {
			number := strings.Join(numbers, ",")
			where := fmt.Sprintf("job_number not in (%s)", number)
			wheres = append(wheres, where)
		}
	}
	for _, c := range awardsStage.AwardsConditionsTableID {
		var where string

		switch c.ConditionSymbol {
		case "", "=":
			where = fmt.Sprintf("%s='%v'", c.ConditionFlag, c.ConditionValue)
		case "!=":
			where = fmt.Sprintf("%s!='%v'", c.ConditionFlag, c.ConditionValue)
		case ">":
			where = fmt.Sprintf("%s>'%v'", c.ConditionFlag, c.ConditionValue)
		case ">=":
			where = fmt.Sprintf("%s>='%v'", c.ConditionFlag, c.ConditionValue)
		case "<":
			where = fmt.Sprintf("%s<'%v'", c.ConditionFlag, c.ConditionValue)
		case "<=":
			where = fmt.Sprintf("%s<='%v'", c.ConditionFlag, c.ConditionValue)
		default:
			where = fmt.Sprintf("%s>='%v' and %s<='%v'",
				c.ConditionFlag, c.ConditionSymbol, c.ConditionFlag, c.ConditionValue)
		}
		wheres = append(wheres, where)
	}
	return strings.Join(wheres, " and "), nil
}

// getLotteryStaffsFromPersonnels 从人事资料转换抽奖表
func getLotteryStaffsFromPersonnels(personnels []PersonnelBasicInformation, isItRedundant string) (staffs []LotteryStaffTemporaryTable) {
	for _, personnel := range personnels {
		staff := LotteryStaffTemporaryTable{
			JobNumber:     personnel.JobNumber,
			Name:          personnel.Name,
			Department:    personnel.Department,
			IsItRedundant: isItRedundant,
		}
		staffs = append(staffs, staff)
	}
	return
}

// getLotteryStaffFromTemporary 从待抽奖表转换抽奖表
func getLotteryStaffFromTemporary(staff LotteryStaffTemporaryTable) LotteryStaffQuantityTable {
	return LotteryStaffQuantityTable{
		ID:            staff.ID,
		StageNumber:   staff.StageNumber,
		JobNumber:     staff.JobNumber,
		Name:          staff.Name,
		Department:    staff.Department,
		IsItRedundant: staff.IsItRedundant,
	}
}

// getSpecifyLotteryStaffs 获取当前阶段指定数量抽奖人员
func getSpecifyLotteryStaffs(lotteryStaffs []LotteryStaffTemporaryTable, quantity int, isItRedundant string) (staffs []LotteryStaffQuantityTable, err error) {
	length := len(lotteryStaffs)
	if length == 0 {
		return
	}
	if length == 1 {
		staff := getLotteryStaffFromTemporary(lotteryStaffs[0])

		err := delSpecifyLotteryStaff(staff.JobNumber, isItRedundant)
		if err != nil {
			return nil, err
		}
		return append(staffs, staff), nil
	}
	for i := 0; i < quantity; i++ {
		index := lotteryStaffIndex(lotteryStaffs, staffs)
		if index == -1 {
			continue
		}
		staff := getLotteryStaffFromTemporary(lotteryStaffs[index])

		err := delSpecifyLotteryStaff(staff.JobNumber, isItRedundant)
		if err != nil {
			return nil, err
		}
		staffs = append(staffs, staff)
		lotteryStaffs = append(lotteryStaffs[:index], lotteryStaffs[index+1:]...)
	}
	return
}

// lotteryStaffIndex 获取抽奖人员唯一索引
func lotteryStaffIndex(lotteryStaffs []LotteryStaffTemporaryTable, staffs []LotteryStaffQuantityTable) int {
	length := len(lotteryStaffs)
	if length == 0 {
		return -1
	}
	if length == 1 {
		return 0
	}
	index := rand.Intn(length - 1)
	current := lotteryStaffs[index]

	for _, s := range staffs {
		if s.JobNumber == current.JobNumber {
			return lotteryStaffIndex(lotteryStaffs, staffs)
		}
	}
	return index
}

// delSpecifyLotteryStaff 从待抽奖人员删除指定人员
func delSpecifyLotteryStaff(job_number, isItRedundant string) error {
	if isItRedundant == "否" {
		err := NewLotteryStaffTemporaryTable().DeleteByCondition("job_number=?", job_number)
		if err != nil {
			return fmt.Errorf("删除符合条件抽奖人员失败, %s", err)
		}
	}
	return nil
}

// Delete: 删除奖项表记录
func (table AwardsTable) Delete() error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	var record AwardsTable

	query := db.
		Preload("AwardsStageTableName").
		Where("ID=?", table.ID).
		Find(&record)

	if query.Error != nil {
		return fmt.Errorf("删除奖项表记录失败，%s", query.Error)
	}
	for _, stage := range record.AwardsStageTableName {
		//删除当前抽奖阶段
		err = stage.Delete()
		if err != nil {
			return err
		}
	}
	//删除当前奖项
	del := db.Delete(&table)
	if del.Error != nil {
		return fmt.Errorf("删除奖项表记录失败, %s", del.Error)
	}
	return nil
}

// DeleteByCondition: 根据指定条件删除奖项表记录
func (table AwardsTable) DeleteByCondition(qs string, args ...interface{}) error {
	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	del := db.
		Where(qs, args...).
		Delete(&table)

	if del.Error != nil {
		return fmt.Errorf("删除奖项表记录失败, %s", del.Error)
	}
	return nil
}
