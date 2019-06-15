package kafka

import (
	"awesomeGo/core/conf"
	"testing"
)

func TestConsumer(t *testing.T) {
	conf.LoadConfig()
	Consumer()
}
