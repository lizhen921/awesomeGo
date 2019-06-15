package timer

import (
	"fmt"
	"time"
)

/**
经过d时间执行,某个匿名函数
*/
func TimeAfterFunc(d time.Duration) {
	done := make(chan bool) //当作协程执行完成的信号量
	f := func() {
		fmt.Println("匿名function 执行")
		done <- true
	}
	//timer: 等待d时间，执行函数f
	go time.AfterFunc(d, f)

	b := <-done //等待协程执行完成的信号
	fmt.Println("timer	执行完成：", b)
	close(done)
}

/**
等待d时间
*/
func TimeAfter(d time.Duration) {
	fmt.Println(time.Now())
	time.After(d)
	select {
	case <-time.After(d):
		fmt.Println(time.Now())
	}
}

func TimeNewTimer() {
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()
	fmt.Println(time.Now())
	t := <-timer.C
	fmt.Println(t, " 定时，3秒后执行")
	timer.Reset(5 * time.Second)
	fmt.Println(time.Now())
	t = <-timer.C
	fmt.Println(t, " 重置，等待5s执行")
}

func TimerTicker(d time.Duration) {
	done := make(chan bool)
	timer := time.NewTimer(d)
	defer timer.Stop()
	go func() {

		for {
			select {
			case <-timer.C:
				fmt.Println("timer执行")
				timer.Reset(d)
			case <-done:
				fmt.Println("done")
			}
		}
	}()

	time.Sleep(10 * time.Second)
	done <- true
}
