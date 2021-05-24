package base

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	//标准base64编码
	data := "abcka中gfd*^&&^*fadf"

	sEnc := base32.StdEncoding.EncodeToString([]byte(data))

	fmt.Println(sEnc)

	sDec, _ := base32.StdEncoding.DecodeString(sEnc)

	fmt.Println(string(sDec))

	//兼容base64编码
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))

	fmt.Println(uEnc)

	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
