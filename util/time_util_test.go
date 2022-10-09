package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

const UserInfoUrl = "https://user.api.autohome.com.cn/api/UserInfo/getuserinfo?_appid=dealer_bi&userid=%s&_timestamp=%d&ctoken=%s"

func TestTime(t *testing.T) {

	c := &http.Client{
		Timeout: time.Duration(500) * time.Millisecond,
	}
	timestamp := time.Now().Unix()
	token := fmt.Sprintf("h@d8u7%d", timestamp)
	phoneUrl := fmt.Sprintf(UserInfoUrl, "226163190", timestamp, Md5(token))

	phoneRes, err := c.Get(phoneUrl)
	if err != nil {
		phoneRes, err = c.Get(phoneUrl)
		if err != nil {
		}
	}
	defer phoneRes.Body.Close()

	phoneContent, err := ioutil.ReadAll(phoneRes.Body)
	if err != nil {
	}

	fmt.Printf(string(phoneContent))
}

func TestMD5(t *testing.T) {

	s := []string{"A", "B", "C"}

	k := s[:1]
	fmt.Println(&s[0] == &k[0])

	u := append(s[:1], s[2:]...)
	fmt.Println(&s[0] == &u[0])

	//s := "appkeypm1_appidad_lp_timestamp1663913250000deviceId240255C3DF9A3113B8554C9BA5BE91CBC57640F8appkey"
	//fmt.Println(Md5(s))
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
