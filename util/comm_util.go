package util

import (
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
