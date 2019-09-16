package conversion

import "fmt"

type Person struct {
	Name    string
	Address *struct {
		Street string
		City   string
	}
}

type Animal struct {
	Name string
	Age  string
}

var data *struct {
	Name    string `json:"name"`
	Address *struct {
		Street string `json:"street"`
		City   string `json:"city"`
	} `json:"address"`
}

var person = (*Person)(data) //类型转换

func assertPerson(i interface{}) Person {
	p, ok := i.(Person) //类型断言
	if !ok {
		fmt.Printf("参数i的类型不是Person，而是%T\n", i)
	} else {
		fmt.Printf("转换成功%+v\n", p)
	}
	return p
}

type String string
