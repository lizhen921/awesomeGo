package go1_19

import (
	"fmt"
	"testing"
)

/*
泛型
泛型形参：T
泛型约束/类型形参列表：int | float32 | float64 ，只能用于这三中类型
*/
type MySlice[T int | float32 | float64] []T

// 等同于下列定义
type MySliceInt []int
type MySliceFloat32 []float32
type MySliceFloat64 []float64

// 泛型
func TestT(t *testing.T) {
	//声明时，需要指定类型实参(int)，这个过程叫做实例化
	var mysliceA MySlice[int]
	mysliceA = []int{1, 2, 3}
	a := Double(mysliceA)
	fmt.Println(a)
	mysliceB := []float32{1.1, 2.2, 3.3}
	b := Double(mysliceB)
	fmt.Println(b)
}

// 泛型方法
func Double[T int | float32 | float64](a []T) (res []T) {
	for i := range a {
		a[i] = a[i] + a[i]
	}
	return a
}

// 更复杂的泛型定义
type MyMap[KEY int | string, VALUE int | float64] map[KEY]VALUE

var a MyMap[string, float64] = map[string]float64{
	"jack_score": 9.6,
	"bob_score":  8.4,
}

type AA struct {
	a interface{}
	b any
}

type SliTest struct {
	Name string
}
