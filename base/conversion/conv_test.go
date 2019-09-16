package conversion

import (
	"fmt"
	"testing"
)

func TestAssert(t *testing.T) {
	p := Person{}
	assertPerson(p)
	q := Animal{}
	assertPerson(q)
}

func TestRecover(t *testing.T) {
	var p Person
	var a Animal
	a = Animal(p)
	fmt.Println(a)
}

func TestType(t *testing.T) {
	var str1 String
	str1 = "123"
	fmt.Printf("%T - %v \n", str1, str1)
	var str2 string
	str2 = string(str1)
	fmt.Printf("%T - %v \n", str2, str2)
	str3 := String(str2)
	fmt.Printf("%T - %v \n", str3, str3)
}
