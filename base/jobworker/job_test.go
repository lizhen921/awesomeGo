package jobworker

import "testing"

func TestJob(t *testing.T) {
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
}
