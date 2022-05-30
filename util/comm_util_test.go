package util

import (
	"fmt"
	"testing"
)

func TestHashStringToUint32(t *testing.T) {
	deviceid := "qwedhfyhvhsgagjasdglsagkasdgdasoigajdgaopj"
	fmt.Println(HashStringToUint32(deviceid))
}
