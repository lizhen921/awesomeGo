package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

type MySqlClient struct {
	Host    string
	Port    int
	DB      string
	User    string
	Pwd     string
	pool    *sql.DB
	MaxIdle int
	MaxOpen int
}

func (mc *MySqlClient) Init() (err error) {
	// 构建 DSN 时尤其注意 loc 和 parseTime 正确设置
	// 东八区，允许解析时间字段
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=%s&parseTime=true",
		mc.User,
		mc.Pwd,
		mc.Host,
		mc.Port,
		mc.DB,
		url.QueryEscape("Asia/Shanghai"))

	// Open 全局一个实例只需调用一次
	mc.pool, err = sql.Open("mysql", uri)
	if err != nil {
		return err
	}

	// 使用前 Ping, 确保 DB 连接正常
	if err = mc.pool.Ping(); err != nil {
		return err
	}

	// 设置最大连接数，一定要设置 MaxOpen
	mc.pool.SetMaxIdleConns(mc.MaxIdle)
	mc.pool.SetMaxOpenConns(mc.MaxOpen)
	return nil
}

func (mc *MySqlClient) GetMySqlClient() *sql.DB {
	err := mc.Init()
	if err != nil {
		panic(err)
	}
	return mc.pool
}
