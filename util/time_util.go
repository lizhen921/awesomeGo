package util

var stdOffset int64 = 8 * 3600 * 1e9

/*
 n天内
*/
func IsInNDay(m, nano1, nano2 int64) bool {
	return (nano2+stdOffset)/(86400*1e9)-(nano1+stdOffset)/(86400*1e9) < m
}
