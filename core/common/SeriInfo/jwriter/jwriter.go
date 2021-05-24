package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
)

/*
// SeriInfo ...
type SeriInfo struct {
}

// Reset ...
func (r *SeriInfo) Reset() {}

// String ...
func (r *SeriInfo) String() string { return "" }

// ProtoMessage ...
func (r *SeriInfo) ProtoMessage() {}
*/

// format: ip1:port1,ip2:port2,...
var redishp = flag.String("r", "", "redis proxy")
var ssdbhp = flag.String("s", "", "ssdb proxys")
var expire = flag.String("expire", "259200", "超时时间设置，默认是3天")
var keyPrefix = flag.String("keyprefix", "tt_", "redis中key前缀")

func main() {
	flag.Parse()

	var errNum, errSsdbNum, totleNum float32

	redisClinets := make([]*redis.Conn, 0)
	if redishp != nil {
		for _, hp := range strings.Split(*redishp, ",") {
			if hp == "" {
				continue
			}
			fmt.Printf("redis conf info : host port[%s], expire[%s] keyprefix[%s]\n", hp, *expire, *keyPrefix)
			s, err := redis.Dial("tcp", hp, redis.DialKeepAlive(1*time.Second), redis.DialPassword("warehouse"), redis.DialConnectTimeout(0*time.Second), redis.DialReadTimeout(5*time.Second), redis.DialWriteTimeout(5*time.Second))
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			redisClinets = append(redisClinets, &s)
		}
	}

	ssdbClinets := make([]*redis.Conn, 0)
	if ssdbhp != nil {
		for _, hp := range strings.Split(*ssdbhp, ",") {
			if hp == "" {
				continue
			}
			fmt.Printf("ssdb conf info : host port[%s], expire[%s] keyprefix[%s]\n", hp, *expire, *keyPrefix)
			s, err := redis.DialTimeout("tcp", hp, 0, 5*time.Second, 5*time.Second)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			ssdbClinets = append(ssdbClinets, &s)
		}
	}

	defer func() {
		for idx := range redisClinets {
			(*redisClinets[idx]).Close()
		}
		for idx := range ssdbClinets {
			(*ssdbClinets[idx]).Close()
		}
	}()

	var lineNum int
	reader := bufio.NewReader(os.Stdin)
	redisClinetCnt := len(redisClinets)
	ssdbClinetCnt := len(ssdbClinets)
	if redisClinetCnt <= 0 && ssdbClinetCnt <= 0 {
		panic("redis and ssdb server is empty.")
	}

	expireTime, err := strconv.Atoi(*expire)
	if err != nil {
		fmt.Println("ERROR 超时时间类型转化失败 ", *expire)
		expireTime = 25000
	}

	var line string
	err = nil
	for err != io.EOF {
		if line, err = reader.ReadString('\n'); err != nil {
			continue
		}

		totleNum++
		//cmd =  key \t value
		var lineArr = strings.Split(strings.TrimSpace(line), "\t")
		if len(lineArr) != 2 {
			fmt.Printf("input line lens err[%s]\n", line)
			continue
		}

		//fmt.Println(cmdArr)
		key := *keyPrefix + lineArr[0]
		//TODO 将json string解析到结构体
		pbStruct := &SeriInfo{}
		if err = json.Unmarshal([]byte(lineArr[1]), pbStruct); err != nil {
			fmt.Printf("json Unmarshal err line[%s],err[%s]\n", line, err)
			continue
		}

		//TODO 将结构体序列化
		data, err := proto.Marshal(pbStruct)
		if err != nil {
			fmt.Printf("将结构体序列化出错\n")
			errNum++
			errSsdbNum++
			continue
		}
		//将数据写入到redis之中

		if redisClinetCnt > 0 {
			idx := lineNum % redisClinetCnt
			if _, err = (*redisClinets[idx]).Do("setex", key, expireTime, data); err != nil {
				fmt.Printf("执行redis命令错误 serverid[%d], line[%v]", idx, line)
				errNum++
			}
		}

		if ssdbClinetCnt > 0 {
			idx := lineNum % ssdbClinetCnt
			_, err = (*ssdbClinets[idx]).Do("setex", key, expireTime, data)
			if err != nil {
				fmt.Printf("执行ssdb命令错误 serverid[%d], line[%v]", idx, line)
				errSsdbNum++
			}
		}
	}

	fmt.Printf("\n totle_num[%f] ,err_num[%f], err_ssdb_num[%f]\n", totleNum, errNum, errSsdbNum)
	if errNum > totleNum*0.05 {
		msg := "发现错误行数超过5%，失败重新导入\n"
		fmt.Println(msg)
		panic(errors.New(msg))
	}
}
