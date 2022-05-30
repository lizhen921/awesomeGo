package other

import "testing"

/*
把最近使用的放在链表头部
最久未使用的逐渐到尾部
新元素放在头部
*/

type LRUCache struct {
	cap  int //容量大小 max
	len  int //当前长度
	cMap map[int]*LinkNode
	head *LinkNode
	tail *LinkNode
}

type LinkNode struct {
	key       int
	value     int
	pre, next *LinkNode
}

func Constructor(cap int) LRUCache {
	return LRUCache{
		cap:  cap,
		cMap: make(map[int]*LinkNode, cap),
	}
}

func (l *LRUCache) Put(key, value int) {
	if node, ok := l.cMap[key]; ok {
		l.moveToHead(key, node) //旧的，移动
		return
	}
	node := &LinkNode{key: key, value: value}
	l.addToHead(key, node) //新的，添加
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.cMap[key]; ok {
		l.moveToHead(key, node) //旧的，移动
		return node.value
	}
	return -1
}

func (l *LRUCache) moveToHead(key int, node *LinkNode) {
	if node.pre == nil {
		return
	}
	next := node.next
	pre := node.pre
	pre.next = next
	l.addToHead(key, node)
}

func (l *LRUCache) addToHead(key int, node *LinkNode) {
	if l.len == 0 {
		l.head = node
	}

	node.next = l.head
	l.head.pre = node
	l.head = node
	l.len++

	if l.len > l.cap {
		l.removeTail()
	}

}

func (l *LRUCache) removeTail() {
	pre := l.tail.pre
	pre.next = nil
	l.tail = pre
	l.len--
}

func TestLru(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	lRUCache.Get(1)    // 返回 1
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.Get(2)    // 返回 -1 (未找到)
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.Get(1)    // 返回 -1 (未找到)
	lRUCache.Get(3)    // 返回 3
	lRUCache.Get(4)    // 返回 4

}
