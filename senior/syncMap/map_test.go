package syncMap

import (
	"testing"
)

func TestRWMap(t *testing.T) {
	testMap := make(map[int]int)
	for i := 0; i < 1000; i++ {
		go writeMap(testMap, i, i)
		go readMap(testMap, i)
	}
}

func readMap(Map map[int]int, key int) int {
	return Map[key]
}

func writeMap(Map map[int]int, key int, value int) {
	Map[key] = value
}

func TestNewMetuxMap(t *testing.T) {
	mutexMap := newMutexMap(5)
	for i := 0; i < 1000; i++ {
		go mutexMap.writeMap(i, i)
		go mutexMap.readMap(i)
	}
}

func TestSyncMap(t *testing.T) {
	//syncMap := make(sync.Map)
}
