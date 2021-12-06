package substring

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"
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

/*
给你一个字符串s，逐个翻转字符串中的所有单词 。

单词是由非空格字符组成的字符串。s中使用至少一个空格将字符串中的单词分隔开。

请你返回一个翻转s中单词顺序并用单个空格相连的字符串。

说明：

输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
翻转后单词间应当仅用一个空格分隔。
翻转后的字符串中不应包含额外的空格。


示例 1：

输入：s = "the sky is blue"
输出："blue is sky the"
示例 2：

输入：s = "hello world "
输出："world hello"
解释：输入字符串可以在前面或者后面包含多余的空格，但是翻转后的字符不能包括。
*/
func reverseWords(s string) string {
	res := ""
	strB := []byte(s)
	reverseString(strB)
	l := len(strB)
	i := 0

	tempB := make([]byte, 0)
	for i < l {
		if strB[i] != ' ' {
			tempB = append(tempB, strB[i])
		}

		if (strB[i] == ' ' || i == l-1) && len(tempB) > 0 {
			reverseString(tempB)
			res = res + " " + string(tempB)
			i++
			tempB = make([]byte, 0)
			continue
		}
		i++
	}
	return res[1:]
}

func TestA(t *testing.T) {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))

}

/*
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
请定义一个函数实现字符串左旋转操作的功能。
比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

示例 1：

输入: s = "abcdefg", k = 2
输出: "cdefgab"

示例 2：

输入: s = "lrlose  umgh", k = 6
输出: "umghlrlose"

*/
func reverseLeftWords(s string, n int) string {
	strB := []byte(s)
	reverseString(strB)

	b1 := strB[:len(strB)-n]
	b2 := strB[len(strB)-n:]
	reverseString(b1)
	reverseString(b2)
	return string(strB)
}

func reverseLeftWords1(s string, n int) string {
	strB := make([]byte, len(s))
	time.Sleep(time.Millisecond * 10)
	for i := 0; i < len(s); i++ {
		if i < n {
			strB[len(s)-n+i] = s[i]
		} else {
			strB[i-n] = s[i]
		}
	}
	return string(strB)
}

func TestReverseLeftWords(t *testing.T) {

	fmt.Println(reverseLeftWords1("lrloseumgh", 6))
}

/*
实现 strStr() 函数。

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

说明：

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

示例 1：

输入：haystack = "hello", needle = "ll"
输出：2
示例 2：

输入：haystack = "aaaaa", needle = "bba"
输出：-1
示例 3：

输入：haystack = "", needle = ""
输出：0
*/

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	if needle == "" {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		j := 0
		start := i
		for j < len(needle) && start < len(haystack) && haystack[start] == needle[j] {
			start++
			j++
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

//寻找next数组，next的含义是
func getNext(next []int, str string) {
	j := 0
	next[0] = j

	for i := 1; i < len(str); i++ {
		for str[i] != str[j] && j > 0 {
			j = next[j-1]
		}
		if str[i] == str[j] {
			j++
		}
		next[i] = j
	}

}

func strStr2(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := make([]int, n)
	getNext(next, needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 回退到j的前一位
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

/*
重复的子字符串

输入: "abcabcabcabc"

输出: True

解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)

aabaab


*/

func repeatedSubstringPattern(s string) bool {
	if len(s) <= 0 {
		return false
	}

}
