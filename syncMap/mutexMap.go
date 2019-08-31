package syncMap

import "sync"

type MutexMap struct {
	sync.RWMutex
	Map map[int]int
}

func newMutexMap(size int) *MutexMap {
	sm := new(MutexMap)
	sm.Map = make(map[int]int, size)
	return sm

}

func (sm *MutexMap) readMap(key int) int {
	sm.RLock()
	value := sm.Map[key]
	sm.RUnlock()
	return value
}

func (sm *MutexMap) writeMap(key int, value int) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}
