package _interface

import (
	"fmt"
	"testing"
)

type People interface {
	Speak(s string)string
}

type Father struct {

}

func (f *Father)Speak(s string) string {

	return "I love you, "+ s
}

func TestInterface(t *testing.T) {
	var peo People = &Father{}
	fmt.Println(peo.Speak("son"))

	var fa *People
	fmt.Println(fa == nil)
	fa2 := getFa()
	fmt.Println(fa2 == nil)
	
}
func getFa() People {
	var fa *Father
	return fa
}