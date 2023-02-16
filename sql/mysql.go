package sql

import (
	"cpic/util"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	UserName     string = "root"
	Password     string = "123456"
	Addr         string = "127.0.0.1"
	Port         int    = 3307
	Database     string = "crawler"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func Connect() *sql.DB {
	//組合sql連線字串
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, Addr, Port, Database)
	//連接MySQL
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
	}
	DB.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	DB.SetMaxOpenConns(MaxOpenConns)
	DB.SetMaxIdleConns(MaxIdleConns)
	return DB
}

// 插入
func Query(sql string) *sql.Rows {
	db := Connect()
	//插入資料
	rows, err := db.Query(sql)
	checkErr(err)
	return rows
}

// 大量寫入
func Trans(sqlDatas []string) bool {
	db := Connect()
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return false
	}
	for _, sch := range sqlDatas {
		_, err = tx.Exec(sch)
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			return false
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// 章節內容寫入到資料庫
func GenInsert(table string, fields []string, values []string) string {
	var sqltxt string
	time.Sleep(300 * time.Millisecond)
	sqltxt = util.CollectInsert(table, fields, values)
	stmt := ExecSql(sqltxt)
	lstId, err := stmt.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	s := strconv.Itoa(int(lstId))
	return s
}

// 插入「
func ExecSql(sqltxt string) sql.Result {
	db := Connect()
	stmt, err := db.Exec(sqltxt)
	checkErr(err)
	return stmt
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
