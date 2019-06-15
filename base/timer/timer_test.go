package timer

import (
	"testing"
	"time"
)

func TestTimerAfterFunc(t *testing.T) {
	TimeAfterFunc(3 * time.Second)
}

func TestTimeAfter(t *testing.T) {
	TimeAfter(3 * time.Second)
}

func TestNewTimer(t *testing.T) {
	TimeNewTimer()
}

func TestTimerTicker(t *testing.T) {
	TimerTicker(3 * time.Second)
}
