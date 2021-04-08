package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

//DBInfo 连接参数
type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	PassWord string
	Charset  string
}

//TableColumn 表基信息
type TableColumn struct {
	ColumnName    string
	DateType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

var DBTypeToStructType = map[string]string{
	"int":       "int32",
	"tinyint":   "int8",
	"smallint":  "int",
	"mediumint": "int64",
	"bigint":    "int64",
	"bit":       "int",
	"bool":      "bool",
	"enum":      "string",
	"set":       "string",
	"varchar":   "string",
	"char":      "string",
	//金钱类型
	//时间类型
}

//NewDBModel 初始化
func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

//Conn 连接数据库
func (m *DBModel) Conn() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?charset=%&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.UserName,
		m.DBInfo.PassWord,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) GetColumns(dbName, TableName string) ([]*TableColumn, error) {
	query := "select column_name,data_type,column_key,is_nullable,column_type,column_comment from columns where table_schema=? and table_name=?"
	rows, err := m.DBEngine.Query(query, dbName, TableName)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("无数据")
	}
	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DateType, &column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil
}
