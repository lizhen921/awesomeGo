package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"testing"
	"time"
)

func TestRedisDemo(t *testing.T) {
	//使用redis封装的Dial进行tcp连接
	//Dial方法有可选参数：DialOption，具体使用可以参见源码conn.go,与redis.DialConnectTimeout(10*time.Second)类似
	conn, err := redis.Dial("tcp", ":6379", redis.DialConnectTimeout(10*time.Second))
	errCheck(err)

	defer conn.Close()

	//对本次连接进行set操作
	_, setErr := conn.Do("set", "url", "哈哈哈哈 github.com/gomodule/redigo/redis")
	errCheck(setErr)

	//使用redis的string类型获取set的k/v信息
	r, getErr := redis.String(conn.Do("get", "url"))
	errCheck(getErr)
	fmt.Println(r)
}

func TestTedisPubSub(t *testing.T) {
	conn, err := redis.Dial("tcp", ":6379", redis.DialConnectTimeout(10*time.Second))
	errCheck(err)

	defer conn.Close()

}

func TestRedisPool(t *testing.T) {
	conn := RedisPool.Get()
	r, getErr := redis.String(conn.Do("get", "url"))
	errCheck(getErr)
	fmt.Println(r)
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

func errCheck(err error) {
	if err != nil {
		fmt.Println("sorry,has some error:", err)
		os.Exit(-1)
	}
}
