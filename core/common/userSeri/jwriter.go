package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

import "github.com/garyburd/redigo/redis"
import "flag"
import "bufio"
import "strings"
import "encoding/json"
import "github.com/golang/protobuf/proto"

func main() {
	var errNum int = 0
	var errSsdbNum int = 0
	var totleNum int = 0

	// format: ip1:port1,ip2:port2,...
	var redishp = flag.String("r", "", "redis proxy")
	var ssdbhp = flag.String("s", "", "ssdb proxys")
	var expire = flag.String("expire", "259200", "超时时间设置，默认是3天")
	var key_prefix = flag.String("keyprefix", "tt_", "redis中key前缀")
	flag.Parse()

	redis_clinets := make([]*redis.Conn, 0)
	if redishp != nil {
		var rediss = strings.Split(*redishp, ",")
		for _, hp := range rediss {
			if hp == "" {
				continue
			}
			fmt.Printf("redis conf info : host port[%s], expire[%s] keyprefix[%s]\n", hp, *expire, *key_prefix)
			// s, err := redis.DialTimeout("tcp", hp, 0, 5*time.Second, 5*time.Second); if err != nil {
			s, err := redis.Dial("tcp", hp, redis.DialKeepAlive(1*time.Second), redis.DialPassword("rcmrealtime"), redis.DialConnectTimeout(0*time.Second), redis.DialReadTimeout(5*time.Second), redis.DialWriteTimeout(5*time.Second))
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			redis_clinets = append(redis_clinets, &s)
		}
	}

	ssdb_clinets := make([]*redis.Conn, 0)
	if ssdbhp != nil {
		var ssdbs = strings.Split(*ssdbhp, ",")
		for _, hp := range ssdbs {
			if hp == "" {
				continue
			}
			fmt.Printf("ssdb conf info : host port[%s], expire[%s] keyprefix[%s]\n", hp, *expire, *key_prefix)
			// s, err := redis.DialTimeout("tcp", hp, 0, 5*time.Second, 5*time.Second); if err != nil {
			s, err := redis.Dial("tcp", hp, redis.DialKeepAlive(1*time.Second), redis.DialPassword("rcmrealtime"), redis.DialConnectTimeout(0*time.Second), redis.DialReadTimeout(5*time.Second), redis.DialWriteTimeout(5*time.Second))
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			ssdb_clinets = append(ssdb_clinets, &s)
		}
	}
	defer func() {
		for idx, _ := range redis_clinets {
			(*redis_clinets[idx]).Close()
		}
		for idx, _ := range ssdb_clinets {
			(*ssdb_clinets[idx]).Close()
		}
	}()
	fmt.Println("connection done")

	var line_num = 0
	reader := bufio.NewReader(os.Stdin)
	redis_clinet_cnt := len(redis_clinets)
	ssdb_clinet_cnt := len(ssdb_clinets)
	if redis_clinet_cnt <= 0 && ssdb_clinet_cnt <= 0 {
		panic("redis and ssdb server is empty.")
	}

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		totleNum += 1
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		line = strings.Trim(line, " ")

		var lineArr = strings.Split(line, "\t") //cmd =  key \t value

		if len(lineArr) != 2 {
			msg := fmt.Sprintf("input line lens err[%s]\n", line)
			os.Stderr.WriteString(msg)
			continue
		}
		//fmt.Println(cmdArr)
		key := *key_prefix + lineArr[0]
		value_str := []byte(lineArr[1])
		//TODO 将json string解析到结构体
		pbStruct := &UserSeri{}
		err = json.Unmarshal(value_str, pbStruct)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("json Unmarshal err line[%s],err[%s]\n", line, err))
			continue
		}
		//TODO 将结构体序列化
		data, err := proto.Marshal(pbStruct)
		if err != nil {
			os.Stderr.WriteString("将结构体序列化出错\n")

			errNum += 1
			errSsdbNum += 1
			continue
		}
		//将数据写入到redis之中

		expire_time, err := strconv.Atoi(*expire)
		if err != nil {
			fmt.Println("ERROR 超时时间类型转化失败 ", *expire)
			expire_time = 25000
		}

		if redis_clinet_cnt > 0 {
			idx := line_num % redis_clinet_cnt
			_, err = (*redis_clinets[idx]).Do("setex", key, expire_time, data)
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("执行ssdb命令错误 serverid[%d].", idx))
				os.Stderr.WriteString(fmt.Sprintf("line:[%s]\n", line))
				errNum += 1
			}
		}

		if ssdb_clinet_cnt > 0 {
			idx := line_num % ssdb_clinet_cnt
			_, err = (*ssdb_clinets[idx]).Do("setex", key, expire_time, data)
			if err != nil {
				os.Stderr.WriteString("执行ssdb命令错误.")
				os.Stderr.WriteString(fmt.Sprintf("line:[%s]\n", line))
				errSsdbNum += 1
			}
		}

		line_num += 1
		if line_num%10000 == 0 {
			//反序列化验证数据
			fmt.Print(".")
		}
	}

	os.Stderr.WriteString(fmt.Sprintf("\n read num :[%d] totle_num[%d] ,err_num[%d], err_ssdb_num[%d]\n", line_num, totleNum, errNum, errSsdbNum))
	if errNum*95 > totleNum {
		msg := "发现错误行数超过95%，失败重新导入\n"
		os.Stderr.WriteString(msg)
		panic(errors.New(msg))
	}
}
