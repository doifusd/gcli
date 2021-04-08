package cmd

import (
	"gcli/internal/sql2struct"
	"log"

	"github.com/spf13/cobra"
)

var (
	username  string
	password  string
	host      string
	charset   string
	dbType    string
	dbName    string
	tableName string
)

var sqlCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转换",
	Long:  "sql转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			PassWord: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Conn()
		if err != nil {
			log.Fatalf("dbModel.connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbmodel.GetColumns err: %v", err)
		}
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssenblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}

	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)

	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "数据库帐号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "数据库密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "数据库帐号")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "uft8mb4", "数据库字符集")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "数据库类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "数据库名")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "表名")

}
