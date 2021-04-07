package main

import (
	"database/sql"
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
