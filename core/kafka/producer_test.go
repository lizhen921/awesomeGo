package kafka

import (
	"awesomeGo/core/conf"
	"testing"
)

func TestProducer(t *testing.T) {
	conf.LoadConfig()
	Producer("test kafka ")
}
