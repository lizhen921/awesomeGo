package util

import "time"

var stdOffset int64 = 8 * 3600 * 1e9

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
