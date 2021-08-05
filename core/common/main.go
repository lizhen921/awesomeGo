//func main() {
//	s, err := redis.Dial("tcp", "dataplatform-lf.codis.lfpsg1.in.autohome.com.cn:19021", redis.DialKeepAlive(1*time.Second), redis.DialPassword("warehouse"), redis.DialConnectTimeout(0*time.Second), redis.DialReadTimeout(5*time.Second), redis.DialWriteTimeout(5*time.Second))
//	for i := 0; i < 100; i++ {
//		start := time.Now()
//		if err != nil {
//			fmt.Printf("Dial err : %v\n", err)
//		}
//		re, err := s.Do("GET", "rts_8355BC021536AC2A31913AC661D33FFAD051C45B")
//		if err != nil {
//			fmt.Printf("Do get err : %v\n,, re:%v", err, re)
//		}
//		end := time.Since(start)
//		fmt.Println(end)
//
//	}
//}

package main

import (
	"errors"
	"fmt"
	errf "github.com/pkg/errors"
	"sync"
	"time"
)

//
//import (
//	"fmt"
//)
//
//type A struct {
//	s string
//}
//
//func main() {
//	var a *A
//	if check(a) {
//		a, err := generate()
//		fmt.Println(a.s, err)
//	}
//	fmt.Println(a.s)
//}
//func generate() (*A, error) {
//	return &A{s: "b"}, nil
//}
//func check(a *A) bool {
//	return true
//}

func doAllocate(nKB int, wg *sync.WaitGroup) {
	var slice []byte
	for i := 0; i < nKB; i++ {
		t := make([]byte, 1024) // 1KB
		slice = append(slice, t...)
		// 大约会执行 50 秒，方便观察内存增长
		time.Sleep(time.Millisecond)
	}
	wg.Done()
	println("doAllocate done")
}
func doIdleAdd(n int64, wg *sync.WaitGroup) {
	var res int64
	for i := int64(0); i < n; i++ {
		res += i
	}
	wg.Done()
	println("doIdleAdd done")
}
func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU()) // needed before go 1.5
	//wg := new(sync.WaitGroup)
	//wg.Add(2)
	//go doAllocate(50*1024, wg) // 程序运行时最多分配 50MB-100MB 内存, 防止影响宿主机
	//t := int64(math.Pow(10, 11))
	//go doIdleAdd(t, wg) // 执行加法，大约会执行 30 秒，方便观察运行情况
	//wg.Wait()

	//URL := "http://lf-summary-sort-prod.autohome.com.cn/api/v1/summary/resourceProbuf?version=1.0&app=1&flag=2&rTypeReq=true&itemKeys=010140080040021-6215513"
	//c := &http.Client{}
	//resp, err := c.Get(URL)
	//fmt.Println(err)
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(err)
	//fmt.Println(string(body))
	//node := &pbsummary.SummaryResponse{}
	//err = proto.Unmarshal(body, node)
	//fmt.Println(err)
	//b, err := json.Marshal(node)
	//fmt.Println(string(b))

	//实验解析
	/*	pvdata := "OAQAAAgAAAACISAAAAgAgACBwEcQhAeU7wcHEAAAgEYABBAoUAUiAIAAEAAAAIIAQAAAAIAAAACAgkIAAAgIYoQBAAAwCAIAACACAwAAIAAABAAQAQBAIIIAAAAQAAAAAAAAACEQEQAYBQAAGFCGgAAASEAFMEGICCCyKgCAKhgAFIAAAAAAAABCAAAAAAQCAAAAAAAAAEAAgAAAACAAIAARSCEAEAACABACAFAgBAAAAIAAEIAAAQIAAAA="
		decoded, _ := base64.StdEncoding.DecodeString(pvdata)
		idxMap := make(map[int]bool)
		for i, v := range decoded {
			if v == 0 {
				continue
			}
			fmt.Printf("%8b:", v)
			for j := 0; j < 8; j++ {
				if v == 0 {
					break
				}
				if (v & 1) == 1 {
					idx := i*8 + j + 1
					idxMap[idx] = true
					fmt.Print(idx, ",")
				}
				v = v >> 1
			}
			fmt.Println("")
		}
		fmt.Println(idxMap)
	*/
	err := errors.New("1")

	fmt.Printf("err EOF : %+v\n", errf.Wrap(err, "2"))
	fmt.Println(errf.WithMessage(err, "3"))
	fmt.Println(errf.WithStack(err))
}

type A struct {
	Number float64 `json:"number"`
}
