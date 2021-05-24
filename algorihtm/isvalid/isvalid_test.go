package isvalid

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	is := isValid("()")
	fmt.Println(is)
}
