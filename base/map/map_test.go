package _map

import (
	"fmt"
	"testing"
)

type Stu struct {
	Name string
}
func TestMapStruct(t *testing.T)  {
	stu1 := Stu{
		"张三",
	}
	stu1.Name = "李四"
	map1 := make(map[string]Stu)
	map1["stu1"] = stu1
	fmt.Println(stu1)
	fmt.Println(map1["stu1"])
	stu1.Name = "王五"

	fmt.Println(stu1)
	fmt.Println(map1["stu1"])
	//map1["stu1"].Name = "lisi" //编译报错
	//map1["stu1"].Name是值引用，无法赋值。这里如果要修改map，则需要定重新对map赋值 如：map1["stu1"] = stu1
	//或者将map定义为map[string]*Stu类型，那么map1["stu1"]便成了指针，修改其引用对象的Name是可以的。
}

type student struct {
	Name string
	Age  int
}

func TestMapFor(t *testing.T) {
	//定义map
	m := make(map[string]*student)

	//定义student数组
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	//将数组依次添加到map中
	for _, stu := range stus {
		fmt.Println(stu.Name, stu)
		m[stu.Name] = &stu
		fmt.Printf("%p\n",&stu) //stu是一个值拷贝，而&stu这个是一个指针，可以看到每次循环进来，该指针是同一个
	}
	fmt.Println(m["zhou"])
	//打印map
	for k,v := range m {
		fmt.Println(k ,"=>", v.Name)
	}
}