package timer

import (
	"testing"
	"time"
)

/**
两种ticker的使用方法
*/

func TestNewTickerFunc(t *testing.T) {
	NewTickerFunc(3 * time.Second)
}

func TestTickFunc(t *testing.T) {
	TickFunc(3 * time.Second)
}
