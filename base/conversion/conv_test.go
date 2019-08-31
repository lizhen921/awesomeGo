package conversion

import (
	"testing"
)

func TestAssert(t *testing.T) {
	p := Person{}
	assertPerson(p)
	q := Animal{}
	assertPerson(q)
}

func TestRecover(t *testing.T) {
	//
	//var p Person
	//var a Animal
	//
	//a = p.(Animal)
	//
	//fmt.Println(a)
}
