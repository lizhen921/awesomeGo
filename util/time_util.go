package util

import (
	"fmt"
	uuid "github.com/rs/xid"
	"math/rand"
	"time"
)

var stdOffset int64 = 8 * 3600 * 1e9

func TimestamToDateStr(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 ")
}

//时间错转时间串
func TimestampToTimeStr(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("15:04:05")
}

/*
 n天内
*/
func IsInNDay(m, nano1, nano2 int64) bool {
	return (nano2+stdOffset)/(86400*1e9)-(nano1+stdOffset)/(86400*1e9) < m
}

//nano为秒级时间戳
func IsToday(nano int64) bool {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
	return (nano-startTime)/86400 < 1 && (nano-startTime) >= 0
}

func IsMillToday(nano int64) bool {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).UnixNano() / 1e6
	fmt.Println(startTime)
	return (nano-startTime)/1e3/86400 < 1 && (nano-startTime) >= 0
}

//生成loguuid
func RandString32() string {
	const kAlphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	nanoStr := fmt.Sprintf("%d", time.Now().UnixNano())
	if len(nanoStr) > 32 {
		return nanoStr[0:32]
	}
	b := make([]byte, 32-len(nanoStr))
	for i := range b {
		b[i] = kAlphaNum[rand.Int63()%int64(len(kAlphaNum))]
	}
	return nanoStr + string(b)
}

func UUID() string {
	return uuid.New().String()
}
