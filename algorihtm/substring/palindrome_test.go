package substring

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	//longestPalindrome("12345ddddddabcdefedcbaabcdefedcbadddddd654321")
	res := longestPalindrome("9012332109")
	fmt.Println(res)
}
