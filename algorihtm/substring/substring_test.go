package substring

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	testString := "pwwkew"
	n := lengthOfLongestSubstring(testString)
	fmt.Println(n)
}
