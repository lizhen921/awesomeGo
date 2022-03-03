package util

import (
	"math/rand"
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
