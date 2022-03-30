package util

import (
	"encoding/base64"
	"math/rand"
	"net"
	"time"
)

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
