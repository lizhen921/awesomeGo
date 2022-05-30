package util

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"hash/fnv"
	"math/rand"
	"net"
	"strconv"
	"time"
)

//将字符串转为uint32
func HashStringToUint32(str string) uint32 {
	//将str的md5结果存成字符串
	strHex := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	//取前8字节（8字节"FFFFFFFF"，对应4字节整数）
	//strHead8 := strHex[0:8]
	//转成对应的4字节无符号整数
	u64HashInt, err := strconv.ParseUint(strHex, 16, 32)
	if err != nil {
		//转换错误，使用原有取值方式
		return HashStringToInt(str)
	} else {
		//转换成功，使用新值
		return uint32(u64HashInt)
	}
}

func HashStringToInt(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}

//生成[min, max]之间的随机数
func GetRandInt(min, max int) int {
	if min >= max {
		return min
	}
	rand.Seed(int64(time.Now().Nanosecond())) //种子
	var ret = rand.Intn(max + 1 - min)
	return ret + min
}

//获取本机ip
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

//编码，解码
// redisName2idx全集: key:value
// currentExpMap命中的: key:bool
func ExpInfo2BitMapUseRedisConfig(redisName2idx map[string]int, currentExpMap map[string]bool) string {
	numField := len(redisName2idx)
	numByte := numField / 8
	if numField%8 > 0 {
		numByte += 1
	}
	expBitMap := make([]byte, numByte)

	for key, v := range currentExpMap {
		if !v {
			continue
		}

		idx := redisName2idx[key] - 1
		if idx < 0 {
			continue
		}

		byteIdx := idx / 8
		byteOffset := uint(idx % 8)

		bitVal := byte(1 << byteOffset)
		expBitMap[byteIdx] |= bitVal
	}
	return base64.StdEncoding.EncodeToString(expBitMap)
}
