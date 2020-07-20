package cmd

import (
	"github.com/spf13/cobra"
	"gotools/internal/sql2struct"
	"log"
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
	Use:   "sql",
	Short: "sql 转换和处理",
	Long:  "sql 转换和处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var sql2StructCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql 转换",
	Long:  "sql 转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		// 连接数据库
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}
		// 渲染模板
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2StructCmd)
	sql2StructCmdFlags := sql2StructCmd.Flags()
	sql2StructCmdFlags.StringVarP(&username, "username", "", "", "请输入数据库的账号")
	sql2StructCmdFlags.StringVarP(&password, "password", "", "", "请输入数据库的密码")
	sql2StructCmdFlags.StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的HOST")
	sql2StructCmdFlags.StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
	sql2StructCmdFlags.StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例类型")
	sql2StructCmdFlags.StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	sql2StructCmdFlags.StringVarP(&tableName, "table", "", "", "请输入表名称")
}
