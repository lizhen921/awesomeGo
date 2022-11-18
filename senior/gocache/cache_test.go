package gocache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"testing"
	"time"
)

func TestGoCache(t *testing.T) {

	IMCache := cache.New(5*time.Second, time.Second)
	IMCache.OnEvicted(func(k string, v interface{}) {
		t := time.Now().UnixMilli()
		IMCache.Set(k, t, cache.DefaultExpiration)
	})

	a := time.Now().UnixMilli()
	for i := 0; i < 100; i++ {
		time.Sleep(3 * time.Second)
		IMCache.Set("time", a, cache.DefaultExpiration)
		fmt.Println(IMCache.Get("time"))
	}

}
