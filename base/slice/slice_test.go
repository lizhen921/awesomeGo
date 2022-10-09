package slice

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
*
使用截取array的形式初始化一个slice，则slice底层就是该数组，那么slice的len为截取的元素的个数，cap为从截取的位置开始到数组的结尾位置。由于slice和array公用一段内存，则修改slice、array都会相互影响
*/
func TestArrayAndSlice(t *testing.T) {

	var strArray = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var strSlice = strArray[5:6]
	fmt.Println("strArray: ", strArray, "\tlen: ", len(strArray), "\tcap: ", cap(strArray))
	fmt.Println("strSlice: ", strSlice, "\tlen: ", len(strSlice), "\tcap: ", cap(strSlice))

	fmt.Println("修改strSlice的第一个元素为：slice-update")
	strSlice[0] = "slice-update"
	fmt.Println("strArray: ", strArray)
	fmt.Println("strSlice: ", strSlice, "\tlen: ", len(strSlice), "\tcap: ", cap(strSlice))

	fmt.Println("为slice添加一个元素为：slice-add")
	strSlice = append(strSlice, "slice-add")
	fmt.Println("strArray: ", strArray)
	fmt.Println("strSlice: ", strSlice, "\tlen: ", len(strSlice), "\tcap: ", cap(strSlice))

	fmt.Println("修改数组下标为6的元素为：array-update")
	strArray[6] = "array-update"
	fmt.Println("strArray: ", strArray)
	fmt.Println("strSlice: ", strSlice, "\tlen: ", len(strSlice), "\tcap: ", cap(strSlice))

	strSlice = append(strSlice, "slice-add")
	strSlice = append(strSlice, "slice-add")
	strSlice = append(strSlice, "slice-add")
	strSlice = append(strSlice, "slice-add")
	strSlice = append(strSlice, "slice-add")
	fmt.Println("strArray: ", strArray, "\tlen: ", len(strArray), "\tcap: ", cap(strArray))
	fmt.Println("strSlice: ", strSlice, "\tlen: ", len(strSlice), "\tcap: ", cap(strSlice))

}

func TestSlice(t *testing.T) {
	slice1 := make([]int, 0, 0)

	for i := 0; i < 2000; i++ {
		slice1 = append(slice1, i)
		fmt.Println("slice1: len: ", len(slice1), "\tcap: ", cap(slice1))
	}

}
func TestJson(t *testing.T) {
	str1 := ""
	fmt.Println(str1)
	//str2 := "{\\\"list\\\":[{\\\"subject_id\\\":21853,\\\"share_url\\\":\\\"autohome://attention/topicpager?name=全民心动车\\\",\\\"rn\\\":\\\"autohome://rninsidebrowser?url=rn%3A%2F%2FFollow_Find%2FTopicList%3Fname%3D%25E5%2585%25A8%25E6%25B0%2591%25E5%25BF%2583%25E5%258A%25A8%25E8%25BD%25A6\\\",\\\"content\\\":\\\"全民心动车\\\"}]}"
	//fmt.Println(str2)
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(str1), &m)
	fmt.Printf("err: [%v]\n", err)
	fmt.Printf("map: [%v]\n", m)

	var summaryRes map[string]string

	for k, v := range summaryRes {
		fmt.Println(k, v)
	}

	var li []string
	fmt.Println(len(li))
}

func TestSlice2(t *testing.T) {
	list := make([]int, 0, 5)
	liStruct := make([]Person, 0, 5)

	list = append(list, 0)

	liStruct = append(liStruct, Person{Name: "张三", Age: 10})

	updateList(list, liStruct)
	fmt.Printf("%v\n", list)
	fmt.Printf("%v\n", liStruct)

}

func updateList(list []int, liStruct []Person) {
	list[0] = 100
	list = append(list, 1)
	list = append(list, 2)
	list = append(list, 3)
	list = append(list, 4)
	list = append(list, 5)

	liStruct[0].Age = 100
	liStruct = append(liStruct, Person{Name: "张三", Age: 11})
	liStruct = append(liStruct, Person{Name: "张三", Age: 12})
	liStruct = append(liStruct, Person{Name: "张三", Age: 13})
	liStruct = append(liStruct, Person{Name: "张三", Age: 14})
	liStruct = append(liStruct, Person{Name: "张三", Age: 15})

	return
}

type Person struct {
	Name string
	Age  int64
}

func TestPrint(t *testing.T) {
	fmt.Printf("%10s %d\n", "123", 12)
	fmt.Printf("%10s %d\n", "1234567", 12)
}
