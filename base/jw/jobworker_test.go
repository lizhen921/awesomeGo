package jw

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestJobWorker(t *testing.T) {
	disp := NewScheduler()
	go disp.Start()
	go func() {
		count := 0
		for i := 0; i < 1000; i++ {
			job := &Job{
				Args: map[string]interface{}{"req": "job1 - " + strconv.Itoa(count)},
				Func: jobOne,
			}
			disp.addJob(job)
			count++
		}
		close(disp.JobChan)
	}()
	time.Sleep(30 * time.Second)
}

func jobOne(args map[string]interface{}) map[string]interface{} {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Println(args)
	return nil
}
