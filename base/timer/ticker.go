package timer

import (
	"fmt"
	"time"
)

func TickFunc(d time.Duration) {
	done := make(chan bool)
	tiker := time.Tick(d)
	go func() {
		for {
			select {
			case <-time.Tick(d):
				// 这里需要注意的是tick方法创建的定时器可以直接返回一个channel，用起来很方便，但是这种ticker并没有主动关闭的方法，
				// 如果在for循环里直接声明，则会生成过个无法关闭的ticker，所以应该放到for循环外定义
				fmt.Println("错误的定时使用方法。。。")
			case <-tiker:
				fmt.Println("正确的定时使用方法。。。")
			case <-done:
				fmt.Println("定时完成！！！")
			}
		}
	}()
	time.Sleep(10 * time.Second)
	done <- true
}

//NewTicker创建定时
func NewTickerFunc(d time.Duration) {
	done := make(chan bool)
	go func() {
		ticker := time.NewTicker(d)
		fmt.Println(time.Now())
		time.Sleep(7 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case t := <-ticker.C:
				fmt.Println(t, "ticker channel time")
				fmt.Println(time.Now(), "定时执行 start。。。")
				//time.Sleep(6 * time.Second)
				fmt.Println(time.Now(), "定时执行 end  。。。")

			case <-done:
				fmt.Println(time.Now(), "定时完成！！！")
			}
		}
	}()
	time.Sleep(10 * time.Second)
	done <- true
}
