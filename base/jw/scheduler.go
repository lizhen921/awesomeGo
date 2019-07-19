package jw

import (
	"fmt"
)

const (
	MAXJOB    = 10000 //最大job缓存个数
	MAXWORKER = 1000  //最大worker数量
)

type Job struct {
	Id   int64
	Name string
	Args map[string]interface{}
	Res  map[string]interface{}
	Func func(map[string]interface{}) map[string]interface{}
}

type Worker struct {
	Id      int
	JobPool chan Job //当前work执行的job队列
	IsBusy  bool     //当前worker是否忙碌
}

type Scheduler struct {
	IdleWorkChan  chan *Worker //空闲的worker队列
	WorkerTimeout int          //取Worker超时时间
	JobChan       chan *Job    //需要处理的Job队列
	JobTimeout    int          //等待job
	MaxWorkers    int          //最大worker数量
}

/**
新建调度器
*/
func NewScheduler() *Scheduler {
	idleWorkChan := make(chan *Worker, MAXWORKER)
	jobChan := make(chan *Job, MAXJOB)
	scheduler := &Scheduler{IdleWorkChan: idleWorkChan, JobChan: jobChan, MaxWorkers: MAXWORKER}
	return scheduler
}

func (d *Scheduler) Start() {
	//创建worker组
	for i := 0; i < d.MaxWorkers; i++ {
		worker := &Worker{
			Id:      i,
			JobPool: make(chan Job),
			IsBusy:  false,
		}
		d.IdleWorkChan <- worker
	}
	d.run()
}

//scheduler分配空闲worker处理job
func (d *Scheduler) run() {
	for {
		select {
		//从worker池中取空闲的worker
		case worker := <-d.IdleWorkChan:
			/*for job := range d.JobChan {
				fmt.Println("---------------------------------------------", job)
				worker.IsBusy = true
				go d.runJob(worker, job)
			}*/
			select {
			//从jobchan中取待处理的job
			case job, ok := <-d.JobChan:
				//判断channel是否关闭，select这种方式读取channel会在已闭关的chanel中读到一次nil
				if ok {
					worker.IsBusy = true
					go func() {
						res := d.runJob(worker, job)
						fmt.Println("todo 处理结果", res)
					}()
				} else {
					fmt.Println("channel 已关闭")
				}
			}
		}
	}
}

func (d *Scheduler) runJob(w *Worker, j *Job) map[string]interface{} {
	fmt.Println("work id : ", w, ", do job :", j)
	j.Res = j.Func(j.Args)
	w.IsBusy = false
	d.IdleWorkChan <- w
	return nil
}

//往Jobchan中添加job
func (d *Scheduler) addJob(job *Job) {
	d.JobChan <- job
}
