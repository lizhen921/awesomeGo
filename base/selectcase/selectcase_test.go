package selectcase

import (
	"fmt"
	"testing"
)

func TestSelectCase(t *testing.T) {

	ch1 := make(chan int, 1000)
	ch2 := make(chan int, 1000)
	for i := 0; i < 100; i++ {
		ch1 <- i
		ch2 <- i
	}
	for i := 0; i < 210; i++ {
		select {
		case c1 := <-ch1:
			fmt.Println("c1:", c1)
		case c2 := <-ch2:
			fmt.Println("c2:", c2)
		default:
			fmt.Println("default:")
		}
	}

}
