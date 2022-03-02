package dp

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{4, 21, 3, 9, 2, 4, 5}
	profit := maxProfit(prices)
	fmt.Println(profit)
}
