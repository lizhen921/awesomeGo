## 1. Map并发读写会有异常：test: TestRWMap
    
    fatal error: concurrent map read and map write

test方法，见: TestRWMap


## 2. 自定义线程安全的map

    见： mutexMap.go


## 3. go1.9以后提供，sync.Map

1. 空间换时间。 通过冗余的两个数据结构(read、dirty),实现加锁对性能的影响。只读read，miss的时候读取dirty，miss达到一定数量，同步dirty到read。
2. 延迟删除。 删除一个键值只是打标记，只有在提升dirty的时候才清理删除的数据
3. 优先从read读取、更新、删除，因为对read的读取不需要锁

主要字段： `read、mu、dirty、misses`

    type Map struct {
    
        // 当涉及到dirty数据的操作的时候，需要使用这个锁
        mu Mutex
        
        // 一个只读的数据结构，因为只读，所以不会有读写冲突。实际上，实际也会更新这个数据的entries,如果entry是未删除的(unexpunged), 并不需要加锁。如果entry已经被删除了，需要加锁，以便更新dirty数据。
        read atomic.Value // readOnly
        
        // dirty数据包含当前的map包含的entries,它包含最新的entries(包括read中未删除的数据,虽有冗余，但是提升dirty字段为read的时候非常快，不用一个一个的复制，而是直接将这个数据结构作为read字段的一部分),有些数据还可能没有移动到read字段中。
        // 对于dirty的操作需要加锁，因为对它的操作可能会有读写竞争。
        // 当dirty为空的时候， 比如初始化或者刚提升完，下一次的写操作会复制read字段中未删除的数据到这个数据中。
        dirty map[interface{}]*entry
        
        // 当从Map中读取entry的时候，如果read中不包含这个entry,会尝试从dirty中读取，这个时候会将misses加一
        // 当misses累积到 dirty的长度的时候， 就会将dirty提升为read,避免从dirty中miss太多次。因为操作dirty需要加锁。
        misses int
    }

    // readOnly是一个不可变结果，以原子的方式存储在Map.read字段中。
    type readOnly struct {
    	m       map[interface{}]*entry
    	amended bool // true if the dirty map contains some key not in m.
    }

    // read和dirty存储都值类型都是*entry，它包含一个指针ip，指向对应都value值
    type entry struct {
    	p unsafe.Pointer // *interface{}
    }
    
    p有三种值：
    1. nil: entry已被删除了，并且m.dirty为nil
    2. expunged: entry已被删除了，并且m.dirty不为nil，而且这个entry不存在于m.dirty中
    3. 其它： entry是一个正常的值

主要方法有：`Load、Store、Delete、Range`

   **Load**：加载方法，也就是提供一个键key,查找对应的值value,如果不存在,通过ok反映
```$xslt
func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
	// 1. 首先从m.read中得到只读readOnly,从它的map中查找，不需要加锁
	read, _ := m.read.Load().(readOnly)
	e, ok := read.m[key]
	// 2. 如果没找到，并且m.dirty中有新数据，需要从m.dirty查找，这个时候需要加锁
	if !ok && read.amended {
		m.mu.Lock()
		// 双检查，避免加锁的时候m.dirty提升为m.read,这个时候m.read可能被替换了。
		read, _ = m.read.Load().(readOnly)
		e, ok = read.m[key]
		// 如果m.read中还是不存在，并且m.dirty中有新数据
		if !ok && read.amended {
			// 从m.dirty查找
			e, ok = m.dirty[key]
			// 不管m.dirty中存不存在，都将misses计数加一
			// missLocked()中满足条件后就会提升m.dirty
			m.missLocked()
		}
		m.mu.Unlock()
	}
	if !ok {
		return nil, false
	}
	return e.load()
}
```
   看下m.dirty的内容什么时候放到readOnly里的：
```$xslt
func (m *Map) missLocked() {
	m.misses++
	if m.misses < len(m.dirty) {
		return
	}
	m.read.Store(readOnly{m: m.dirty})
	m.dirty = nil
	m.misses = 0
}
```
**Store**：更新或者新增一个entry
```$xslt
func (m *Map) Store(key, value interface{}) {
	// 如果m.read存在这个键，并且这个entry没有被标记删除，尝试直接存储。
	// 因为m.dirty也指向这个entry,所以m.dirty也保持最新的entry。
	read, _ := m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok && e.tryStore(&value) {
		return
	}
	// 如果`m.read`不存在或者已经被标记删除
	m.mu.Lock()
	read, _ = m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok {
		if e.unexpungeLocked() { //标记成未被删除
			m.dirty[key] = e //m.dirty中不存在这个键，所以加入m.dirty
		}
		e.storeLocked(&value) //更新
	} else if e, ok := m.dirty[key]; ok { // m.dirty存在这个键，更新
		e.storeLocked(&value)
	} else { //新键值
		if !read.amended { //m.dirty中没有新的数据，往m.dirty中增加第一个新键
			m.dirtyLocked() //从m.read中复制未删除的数据
			m.read.Store(readOnly{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry(value) //将这个entry加入到m.dirty中
	}
	m.mu.Unlock()
}
func (m *Map) dirtyLocked() {
	if m.dirty != nil {
		return
	}
	read, _ := m.read.Load().(readOnly)
	m.dirty = make(map[interface{}]*entry, len(read.m))
	for k, e := range read.m {
		if !e.tryExpungeLocked() {
			m.dirty[k] = e
		}
	}
}
func (e *entry) tryExpungeLocked() (isExpunged bool) {
	p := atomic.LoadPointer(&e.p)
	for p == nil {
		// 将已经删除标记为nil的数据标记为expunged
		if atomic.CompareAndSwapPointer(&e.p, nil, expunged) {
			return true
		}
		p = atomic.LoadPointer(&e.p)
	}
	return p == expunged
}
```


使用方法：``







