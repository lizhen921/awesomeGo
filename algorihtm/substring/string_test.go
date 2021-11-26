package substring

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

输入: s = "anagram", t = "nagaram"
输出: true

输入: s = "rat", t = "car"
输出: false
*/

func isAnagram(s string, t string) bool {
	sliceS := strings.Split(s, "")
	sliceT := strings.Split(t, "")
	sort.Strings(sliceS)
	sort.Strings(sliceT)
	s = strings.Join(sliceS, "")
	t = strings.Join(sliceT, "")
	if s == t {
		return true
	}
	return false
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	strMap := make(map[rune]int)
	for _, v := range s {
		strMap[v]++
	}

	for _, v := range t {
		strMap[v]--
		if strMap[v] < 0 {
			return false
		}
	}
	return true
}

/*
反转字符串
*/
func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Printf("%s\n", s)
	reverseString(s)
	fmt.Printf("%s", s)
}

/*
反转字符串 II
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

示例 1：

输入：s = "abcdefg", k = 2
输出："bacdfeg"
示例 2：

输入：s = "abcd", k = 2
输出："bacd"

*/
func reverseStr(s string, k int) string {
	sByte := []byte(s)

	l := 0
	for l < len(sByte) {
		i := l
		j := l + k - 1
		if j > len(sByte)-1 {
			j = len(sByte) - 1
		}
		for i < j {
			sByte[i], sByte[j] = sByte[j], sByte[i]
			i++
			j--
		}
		l = l + 2*k
	}
	return string(sByte)
}

/*
替换字符串   %20
*/
func replaceSpace(s string) string {
	count := 0
	for _, v := range s {
		if v == ' ' {
			count++
		}
	}
	if count == 0 {
		return s
	}
	l := 2 * count
	temp := make([]byte, l)
	b := []byte(s)
	b = append(b, temp...)

	sIndex := len(s) - 1
	for i := len(b) - 1; i > 0; {
		if s[sIndex] == ' ' {
			b[i] = '0'
			b[i-1] = '2'
			b[i-2] = '%'
			i -= 3
			sIndex--
		} else {
			b[i] = s[sIndex]
			i--
			sIndex--
		}

	}
	return string(b)
}
