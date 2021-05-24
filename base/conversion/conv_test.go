package conversion

import (
	"fmt"
	"testing"
	"time"
)

type AnimalDQ []Animal

func TestAssert(t *testing.T) {
	//p := Person{}
	//assertPerson(p)
	//q := Animal{}
	//assertPerson(q)
	var m map[string]Animal
	var a1 Animal
	var a2 Animal
	var aList AnimalDQ
	fmt.Printf("%v\n", aList)
	fmt.Printf("%v\n", aList == nil)

	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", m == nil)

	aList = append(aList, a1)
	aList = append(aList, a2)
	fmt.Printf("%v\n", aList)
	fmt.Printf("%v\n", aList[0])
	fmt.Printf("%v\n", aList[0].Age)

	m["1"] = a1
	fmt.Printf("%v\n", m)

}

func TestRecover(t *testing.T) {
	//var p Person
	//var a Animal
	////a = Animal(p)
	//fmt.Println(a)
}

func TestType(t *testing.T) {
	//	var str1 String
	//	str1 = "123"
	//	fmt.Printf("%T - %v \n", str1, str1)
	//	var str2 string
	//	str2 = string(str1)
	//	fmt.Printf("%T - %v \n", str2, str2)
	//	str3 := String(str2)
	//	fmt.Printf("%T - %v \n", str3, str3)
	//
	//	a := []int{0, 1, 2, 3, 4}
	//	//删除第i个元素
	//	i := 2
	//	a = append(a[:i], a[i+1:]...)
	//	fmt.Println(a)
	//}
	timeOutCfg := make(map[string]string)
	timeOutCfg["after_rank"] = "110ms"
	timeDura, err := time.ParseDuration(timeOutCfg["after_rank"])
	fmt.Printf("%v, %v", timeDura, err)

}
