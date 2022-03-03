package designpattern

import "sync"

//单例模式对象类型
type XXXType struct {
}

//单例模式全局对象
var Instance *XXXType

//-------使用sync.Once实现-------
var once sync.Once

func GetInstance() *XXXType {
	once.Do(func() {
		Instance = new(XXXType)
	})
	return Instance
}

//-------使用sync.Once实现done-------

//-------使用Lock实现-------
var lock sync.Mutex

func GetInstance2() *XXXType {
	if Instance == nil {
		lock.Lock()
		if Instance == nil {
			Instance = new(XXXType)
		}
		lock.Unlock()
	}
	return Instance
}

//-------使用Lock实现 done-------
