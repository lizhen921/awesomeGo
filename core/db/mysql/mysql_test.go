package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func getMySqlClient() *sql.DB {
	mysqlClient := MySqlClient{
		Host:    "127.0.0.1",
		Port:    3306,
		DB:      "awesome_go",
		User:    "root",
		Pwd:     "root",
		MaxIdle: 5,
		MaxOpen: 5,
	}
	return mysqlClient.GetMySqlClient()
}

type TbAwesome struct {
	Id         int
	Title      string
	Author     string
	IsDelete   bool
	CreateUser string
	CreateTime string    //mysql的datetime类型映射成字符串
	UpdateTime int64     //mysql的datetime类型映射成时间戳
	DeleteTime time.Time //mysql都datetime映射成time.Time类型
}

func TestQuery(t *testing.T) {
	db := getMySqlClient()
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	var rows *sql.Rows
	rows, err = db.Query("select id, title, author, id_delete, create_user, create_time, update_time, delete_time from tb_awesome")
	defer rows.Close() //虽然rows.Next()之后会自动关闭rows，但是如果在遍历过程中提前返回，则rows不会自动关闭，所以这里还是要有defer关闭
	if err != nil {
		panic(err)
	}

	tbAwesomeSlice := make([]*TbAwesome, 0, 10)
	for rows.Next() {
		tbAwesome := &TbAwesome{}

		var updateTime time.Time

		rows.Scan(&tbAwesome.Id, &tbAwesome.Title, &tbAwesome.Author, &tbAwesome.IsDelete, &tbAwesome.CreateUser, &tbAwesome.CreateTime, &updateTime, &tbAwesome.DeleteTime)

		tbAwesome.UpdateTime = updateTime.Unix()

		tbAwesomeSlice = append(tbAwesomeSlice, tbAwesome)
	}
	//检查错误，需要确认已经正确的处理了所有行
	if rows.Err() != nil {
		//log.Error("err: ", err)
		fmt.Errorf("err: %v", err)
	}

	fmt.Printf("slice: %+v", tbAwesomeSlice)
}

func TestQueryOne(t *testing.T) {
	db := getMySqlClient()
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	tbAwesome := &TbAwesome{}
	var row *sql.Row
	row = db.QueryRow("select * from tb_awesome where id = 1 ")
	err = row.Scan(&tbAwesome.Id, &tbAwesome.Title, &tbAwesome.Author, &tbAwesome.IsDelete, &tbAwesome.CreateUser, &tbAwesome.CreateTime, &tbAwesome.UpdateTime, &tbAwesome.DeleteTime)
	if err != nil {
		fmt.Errorf("err: %v", err)
	}
	fmt.Printf("tbAwesome: %+v", tbAwesome)
}

func TestExecute(t *testing.T) {
	tbAwesome := &TbAwesome{
		Title:      "文章4",
		Author:     "上官",
		IsDelete:   false,
		CreateUser: "admin",
	}

	db := getMySqlClient()
	db.Exec("")

	//TODO
	fmt.Println(tbAwesome)
}

func TestMysqlDemo(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/awesome_go?charset=utf8&allowOldPasswords=1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from tb_awesome")

	for rows.Next() {
		//row.Scan(...)
	}
	rows.Close()
}
