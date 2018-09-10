package models

import (
	"fmt"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DataDictionary struct {
	ColumnName             string   `gorm:"column:column_name" json:"column_name"`                           //列名
	DataType               string   `gorm:"column:data_type" json:"data_type"`                               //数据类型
	CharacterMaximumLength int      `gorm:"column:character_maximum_length" json:"character_maximum_length"` //字符长度
	NumericPrecision       int      `gorm:"column:numeric_precision" json:"numeric_precision"`               //数字长度
	NumericScale           int      `gorm:"column:numeric_scale" json:"numeric_scale"`                       //小数位数
	IsNullable             string   `gorm:"column:is_nullable" json:"is_nullable"`                           //是否允许非空
	IsAutoIncrement        int      `gorm:"column:is_auto_increment" json:"is_auto_increment"`               //是否自增
	ColumnDefault          string   `gorm:"column:column_default" json:"column_default"`                     //默认值
	ColumnComment          string   `gorm:"column:column_comment" json:"column_comment"`                     //备注
	ColumnValues           []string `json:"column_values"`                                                   //自定义列值
}

func GetDataDictionary(fullname string) ([]*DataDictionary, error) {
	if fullname == "" {
		return nil, fmt.Errorf("获取数据字典失败： 表名为空")
	}
	database, table := getTableSchemaAndName(fullname)

	db, err := gorm.Open(lucky_drawDataSource())
	if err != nil {
		return nil, fmt.Errorf("创建orm引擎失败, %s", err)
	}
	defer db.Close()

	sql := `SELECT column_name,data_type,character_maximum_length,
	numeric_precision,numeric_scale,is_nullable,column_default,column_comment,    
    CASE WHEN extra = 'auto_increment' THEN 1 ELSE 0 END AS 'is_auto_increment'    
    FROM information_schema.columns 
	WHERE table_schema='%s' and table_name='%s';`

	sql = fmt.Sprintf(sql, database, table)

	var data_dictionary []*DataDictionary

	query := db.
		Raw(sql).
		Scan(&data_dictionary)

	if query.Error != nil {
		return nil, fmt.Errorf("获取数据字典失败,%s", query.Error)
	}
	for _, data := range data_dictionary {
		values, err := getColumnValues(db, table, data.ColumnName)
		if err != nil {
			return nil, err
		}
		data.ColumnValues = values
	}
	return data_dictionary, nil
}

func getTableSchemaAndName(fullname string) (db, table string) {
	db_name := strings.SplitN(fullname, ".", 2)
	if len(db_name) == 2 && db_name[0] != "" && db_name[1] != "" {
		db = db_name[0]
		table = db_name[1]
	} else {
		db = lucky_drawDbName()
		table = fullname
	}
	return
}

func getColumnValues(db *gorm.DB, tablename, column string) (results []string, err error) {
	sql := fmt.Sprintf("SELECT DISTINCT %s FROM %s WHERE %s!=''", column, tablename, column)

	rows, err := db.Raw(sql).Rows()
	if err != nil {
		return nil, fmt.Errorf("获取表单数据失败, %s", err)
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("获取表单字段名失败, %s", err)
	}
	scans := make([]interface{}, len(columns))
	values := make([][]byte, len(columns))

	for i := range values {
		scans[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scans...)
		if err != nil {
			return nil, fmt.Errorf("获取表单单个数据失败, %s", err)
		}
		for _, col := range values {
			results = append(results, string(col))
		}
	}
	if len(results) > 10 {
		return nil, nil
	}
	return
}
