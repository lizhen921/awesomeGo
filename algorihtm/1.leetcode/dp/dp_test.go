package dp

import (
	"fmt"
	"github.com/leesper/go_rng"
	"math/rand"
	"testing"
	"time"
)

func TestGen(t *testing.T) {

	dateMonth := time.Now().AddDate(0, 0, 0).Format("1")
	fmt.Println(dateMonth)
	total := 10
	click := 10
	unclick := total - click
	grng := rng.NewBetaGenerator(time.Now().UnixNano())
	//1. 用户的排行榜的统计值 -- redis取  格式
	//2. 车系命中的排行榜及其对应排行。
	/*
		alpha: 某个排行的统计值-点击数
		beta: 某个排行的统计值-未点击数
	*/
	pred := grng.Beta(float64(1+click), float64(1+unclick))

	fmt.Println(pred)

}

/*
"dogs"
[]
*/
func TestSsplitNums(t *testing.T) {

	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(4)
		fmt.Println(r)
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(4)
	fmt.Println(r)
	//a := []int{1, 2, 3, 2, 1}
	//b := []int{3, 2, 1, 4, 7}
	//r := SimpleSlice(a, b)
	//fmt.Println(r)
}
func TestBag(t *testing.T) {
	weight := []int{1, 3, 4, 5}
	value := []int{15, 20, 30, 3}
	res := bag(weight, value, 7)
	fmt.Println(res)
	res = bag2(weight, value, 7)
	fmt.Println(res)
}

func TestNnumTree(t *testing.T) {
	fmt.Println(numTree(3))
}

func TestIintSlpit(t *testing.T) {
	intSlpit(3)
}

func TestFb(t *testing.T) {
	fb(10)
}

func TestMincs(t *testing.T) {
	a := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	mincs(a)
}

func TestDdiffPaths01(t *testing.T) {

	a := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	diffPaths01(a)
}
