package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "fileserver"
)

func init() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/fileserver?charset=utf8")
	//dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	//db ,err := sql.Open("mysql",dsn)
	if err != nil{
		fmt.Printf("Open mysql failed,err:%v\n",err)
		return
	}
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
}

// DBConn : 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}

func ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		//将行数据保存到record字典
		err := rows.Scan(scanArgs...)
		checkErr(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		records = append(records, record)
	}
	return records
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
