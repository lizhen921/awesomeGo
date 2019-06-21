package gc

import "fmt"

/**
逃逸分析
*/

func Sum() int {
	numbers := make([]int, 100)
	for i := range numbers {
		numbers[i] = i + 1
	}
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

const Width, Height = 640, 480

type Coursor struct {
	X, Y int
}

func Center(c *Coursor) {
	c.X = Width / 2
	c.Y = Height / 2
}

func CenterCoursor() {
	c := new(Coursor)
	Center(c)
	fmt.Println(c.X, c.Y)
}
