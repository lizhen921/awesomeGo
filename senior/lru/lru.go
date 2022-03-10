package lru

import (
	"container/list"
	"sync"
)

type Lru struct {
	size  int
	List  list.List
	Cache map[interface{}]*list.Element
	lock  sync.Mutex
}
