# golang Mysql

golang中`database/sql`是操作数据库时经常用的包，这个里面定义了一些sql操作的接口，具体的实现还要根据不同数据库实现具体操作，还有一个重要功能就是连接池。

## simple Demo
```$xslt
package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestMysqlDemo(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/awesome_go?charset=utf8&allowOldPasswords=1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from tb_awesome")

	for rows.Next() {
		//row.Scan(...)
	}
	rows.Close()
}
```
基本操作很简单，先`Open`，然后`Query`，然后遍历结果。这些方法都是`database/sql`包里的接口，具体实现还是在`go-sql-driver/mysql`中。

## 驱动注册：

`import _ "github.com/go-sql-driver/mysql"`这种导包方式，只执行包里的`init()`方法，mysql驱动正是这种方式注册到`database/sql`中的。

```$xslt
func init() {
	sql.Register("mysql", &MySQLDriver{})
}

type MySQLDriver struct{}

func (d MySQLDriver) Open(dsn string) (driver.Conn, error) {
    ...
}
```

`init()`通过调用`Register`方法，将`mysql`驱动对象`MySQLDriver`注册到`database/sql`的全局变量`drivers`中

```$xslt
func Register(name string, driver driver.Driver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("sql: Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("sql: Register called twice for driver " + name)
	}
	drivers[name] = driver
}

// Driver is the interface that must be implemented by a database
// driver.
type Driver interface {
	// Open returns a new connection to the database.
	// The name is a string in a driver-specific format.
	//
	// Open may return a cached connection (one previously
	// closed), but doing so is unnecessary; the sql package
	// maintains a pool of idle connections for efficient re-use.
	//
	// The returned connection is only used by one goroutine at a
	// time.
	Open(name string) (Conn, error)
}
```
`MySQLDriver`实现的就是`Driver`接口，如果我们同时用到多种数据库，就可以通过调用`sql.Register`将不同数据库的实现注册到`sql.drivers`中去，
用的时候再根据注册的`name`将对应的`driver`取出。

## 连接池

![image](http://github.com/altairlee/awesomeGo/blob/master/images/db/go_mysql.jpg?raw=true)

### 初始化DB

`db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/awesome_go?charset=utf8&allowOldPasswords=1")`

实现源码：

```$xslt
func Open(driverName, dataSourceName string) (*DB, error) {
	driversMu.RLock()
	driveri, ok := drivers[driverName]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("sql: unknown driver %q (forgotten import?)", driverName)
	}

	if driverCtx, ok := driveri.(driver.DriverContext); ok {
		connector, err := driverCtx.OpenConnector(dataSourceName)
		if err != nil {
			return nil, err
		}
		return OpenDB(connector), nil
	}

	return OpenDB(dsnConnector{dsn: dataSourceName, driver: driveri}), nil
}
```

`sql.Open`这里只是初始化一个`sql.DB`对象，并没有建立连接。DB对象还是很重要的一个结构

```$xslt
type DB struct {
	// Atomic access only. At top of struct to prevent mis-alignment
	// on 32-bit platforms. Of type time.Duration.
	waitDuration int64 // Total time waited for new connections.

	connector driver.Connector
	// numClosed is an atomic counter which represents a total number of
	// closed connections. Stmt.openStmt checks it before cleaning closed
	// connections in Stmt.css.
	numClosed uint64

	mu           sync.Mutex // protects following fields
	freeConn     []*driverConn   
	connRequests map[uint64]chan connRequest  // 阻塞请求队列，等连接数达到最大限制时，后续请求将插入此队列等待可用连接
	nextRequest  uint64 // Next key to use in connRequests.
	numOpen      int    // number of opened and pending open connections 已建立连接或等待建立连接数
	// Used to signal the need for new connections
	// a goroutine running connectionOpener() reads on this chan and
	// maybeOpenNewConnections sends on the chan (one send per needed connection)
	// It is closed during db.Close(). The close tells the connectionOpener
	// goroutine to exit.
	openerCh          chan struct{}
	resetterCh        chan *driverConn
	closed            bool
	dep               map[finalCloser]depSet
	lastPut           map[*driverConn]string // stacktrace of last conn's put; debug only
	maxIdle           int                    // zero means defaultMaxIdleConns; negative means 0  最大空闲连接数
	maxOpen           int                    // <= 0 means unlimited   数据库最大连接数
	maxLifetime       time.Duration          // maximum amount of time a connection may be reused  连接最长存活期，超过这个时间连接将不再被复用
	cleanerCh         chan struct{}
	waitCount         int64 // Total number of connections waited for.
	maxIdleClosed     int64 // Total number of connections closed due to idle.
	maxLifetimeClosed int64 // Total number of connections closed due to max free limit.

	stop func() // stop cancels the connection opener and the session resetter.
}
```

### 获取连接

真正查询时才会获取连接
`rows, err := db.Query("select * from tb_awesome")`

```$xslt
// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (db *DB) Query(query string, args ...interface{}) (*Rows, error) {
	return db.QueryContext(context.Background(), query, args...)
}

// QueryContext executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	var rows *Rows
	var err error
	for i := 0; i < maxBadConnRetries; i++ {
		rows, err = db.query(ctx, query, args, cachedOrNewConn)
		if err != driver.ErrBadConn {
			break
		}
	}
	if err == driver.ErrBadConn {
		return db.query(ctx, query, args, alwaysNewConn)
	}
	return rows, err
}
```

获取连接的策略有两种
1. cachedOrNewConn(有可用空闲连接则优先使用，没有则创建)
2. alwaysNewConn(不管有没有空闲连接都重新创建)，
```$xslt
const (
	// alwaysNewConn forces a new connection to the database.
	alwaysNewConn connReuseStrategy = iota
	// cachedOrNewConn returns a cached connection, if available, else waits
	// for one to become available (if MaxOpenConns has been reached) or
	// creates a new database connection.
	cachedOrNewConn
)
```
`db.query`中会使用`db.conn`获取连接:
```$xslt
// conn returns a newly-opened or cached *driverConn.
func (db *DB) conn(ctx context.Context, strategy connReuseStrategy) (*driverConn, error) {
	db.mu.Lock()
	if db.closed {
		db.mu.Unlock()
		return nil, errDBClosed
	}
	// Check if the context is expired.
	select {
	default:
	case <-ctx.Done():
		db.mu.Unlock()
		return nil, ctx.Err()
	}
	lifetime := db.maxLifetime

	// Prefer a free connection, if possible.  从freeConn中获取一个空闲连接
	numFree := len(db.freeConn)
	if strategy == cachedOrNewConn && numFree > 0 {
		conn := db.freeConn[0] 
		copy(db.freeConn, db.freeConn[1:])
		db.freeConn = db.freeConn[:numFree-1]
		conn.inUse = true
		db.mu.Unlock()
		if conn.expired(lifetime) {
			conn.Close()
			return nil, driver.ErrBadConn
		}
		// Lock around reading lastErr to ensure the session resetter finished.
		conn.Lock()
		err := conn.lastErr
		conn.Unlock()
		if err == driver.ErrBadConn {
			conn.Close()
			return nil, driver.ErrBadConn
		}
		return conn, nil
	}

    // 如果没有空闲连接，而且当前建立的连接数已经达到最大限制则将请求加入connRequests队列，
    // 并阻塞在这里，直到其它协程将占用的连接释放或connectionOpenner创建
	// Out of free connections or we were asked not to use one. If we're not
	// allowed to open any more connections, make a request and wait.
	if db.maxOpen > 0 && db.numOpen >= db.maxOpen {
		// Make the connRequest channel. It's buffered so that the
		// connectionOpener doesn't block while waiting for the req to be read.
		req := make(chan connRequest, 1)
		reqKey := db.nextRequestKeyLocked()
		db.connRequests[reqKey] = req
		db.waitCount++
		db.mu.Unlock()

		waitStart := time.Now()

		// Timeout the connection request with the context. 请求链接超时
		select {
		case <-ctx.Done():
			// Remove the connection request and ensure no value has been sent
			// on it after removing.
			db.mu.Lock()
			delete(db.connRequests, reqKey)
			db.mu.Unlock()

			atomic.AddInt64(&db.waitDuration, int64(time.Since(waitStart)))

			select {
			default:
			case ret, ok := <-req:
				if ok && ret.conn != nil {
					db.putConn(ret.conn, ret.err, false)
				}
			}
			return nil, ctx.Err()
		case ret, ok := <-req:
			atomic.AddInt64(&db.waitDuration, int64(time.Since(waitStart)))

			if !ok {
				return nil, errDBClosed
			}
			if ret.err == nil && ret.conn.expired(lifetime) {
				ret.conn.Close()
				return nil, driver.ErrBadConn
			}
			if ret.conn == nil {
				return nil, ret.err
			}
			// Lock around reading lastErr to ensure the session resetter finished.
			ret.conn.Lock()
			err := ret.conn.lastErr
			ret.conn.Unlock()
			if err == driver.ErrBadConn {
				ret.conn.Close()
				return nil, driver.ErrBadConn
			}
			return ret.conn, ret.err
		}
	}

	db.numOpen++ // optimistically
	db.mu.Unlock()
	ci, err := db.connector.Connect(ctx)  //调用具体的实现来建立连接
	if err != nil {
		db.mu.Lock()
		db.numOpen-- // correct for earlier optimism
		db.maybeOpenNewConnections()
		db.mu.Unlock()
		return nil, err
	}
	db.mu.Lock()
	dc := &driverConn{
		db:        db,
		createdAt: nowFunc(),
		ci:        ci,
		inUse:     true,
	}
	db.addDepLocked(dc, dc)
	db.mu.Unlock()
	return dc, nil
}
```

总结一下上面获取连接的过程：
1. 首先检查下freeConn里是否有空闲连接，如果有且未超时则直接复用，返回连接，如果没有或连接已经过期则进入下一步；
2. 检查当前已经建立及准备建立的连接数是否已经达到最大值，如果达到最大值也就意味着无法再创建新的连接了，当前请求需要在这等着连接释放，这时当前协程将创建一个channel：
chan connRequest，并将其插入db.connRequests队列，然后阻塞在接收chan connRequest上，等到有连接可用时这里将拿到释放的连接，检查可用后返回或者超时取消获取连接请求；
如果还未达到最大值则进入下一步；
3. 创建一个连接，首先将numOpen加1，然后再创建连接，如果等到创建完连接再把numOpen加1会导致多个协程同时创建连接时一部分会浪费，所以提前将numOpen占住，创建失败再将其减掉；
如果创建连接成功则返回连接，失败则进入下一步
4. 创建连接失败时有一个善后操作，当然并不仅仅是将最初占用的numOpen数减掉，更重要的一个操作是通知connectionOpener协程根据db.connRequests等待的长度创建连接，这个操作的原因是：
numOpen在连接成功创建前就加了1，这时候如果numOpen已经达到最大值再有获取conn的请求将阻塞在step2，这些请求会等着先前进来的请求释放连接，假设先前进来的这些请求创建连接全部失败，
那么如果它们直接返回了那些等待的请求将一直阻塞在那，因为不可能有连接释放(极限值，如果部分创建成功则会有部分释放)，直到新请求进来重新成功创建连接，显然这样是有问题的，
所以maybeOpenNewConnections将通知connectionOpener根据db.connRequests长度及可创建的最大连接数重新创建连接，然后将新创建的连接发给阻塞的请求。

注意：如果maxOpen=0将不会有请求阻塞等待连接，所有请求只要从freeConn中取不到连接就会新创建。

另外Query、Exec有个重试机制，首先优先使用空闲连接，如果2次取到的连接都无效则尝试新创建连接。

获取到可用连接后将调用具体数据库的driver处理sql。

### 释放连接

数据库连接使用完成后会放回到资源池，供其他请求复用，释放连接到操作是`db.putConn`

```$xslt
// putConn adds a connection to the db's free pool.
// err is optionally the last error that occurred on this connection.
func (db *DB) putConn(dc *driverConn, err error, resetSession bool) {
	db.mu.Lock()
	if !dc.inUse {
		if debugGetPut {
			fmt.Printf("putConn(%v) DUPLICATE was: %s\n\nPREVIOUS was: %s", dc, stack(), db.lastPut[dc])
		}
		panic("sql: connection returned that was never out")
	}
	if debugGetPut {
		db.lastPut[dc] = stack()
	}
	dc.inUse = false

	for _, fn := range dc.onPut {
		fn()
	}
	dc.onPut = nil
	//是否无效连接
	if err == driver.ErrBadConn {
		// Don't reuse bad connections.
		// Since the conn is considered bad and is being discarded, treat it
		// as closed. Don't decrement the open count here, finalClose will
		// take care of that.
		db.maybeOpenNewConnections()
		db.mu.Unlock()
		dc.Close()
		return
	}
	if putConnHook != nil {
		putConnHook(db, dc)
	}
	if db.closed {
		// Connections do not need to be reset if they will be closed.
		// Prevents writing to resetterCh after the DB has closed.
		resetSession = false
	}
	if resetSession {
		if _, resetSession = dc.ci.(driver.SessionResetter); resetSession {
			// Lock the driverConn here so it isn't released until
			// the connection is reset.
			// The lock must be taken before the connection is put into
			// the pool to prevent it from being taken out before it is reset.
			dc.Lock()
		}
	}
	added := db.putConnDBLocked(dc, nil)
	db.mu.Unlock()

	if !added {
		if resetSession {
			dc.Unlock()
		}
		dc.Close()
		return
	}
	if !resetSession {
		return
	}
	select {
	default:
		// If the resetterCh is blocking then mark the connection
		// as bad and continue on.
		dc.lastErr = driver.ErrBadConn
		dc.Unlock()
	case db.resetterCh <- dc:
	}
}

// Satisfy a connRequest or put the driverConn in the idle pool and return true
// or return false.
// putConnDBLocked will satisfy a connRequest if there is one, or it will
// return the *driverConn to the freeConn list if err == nil and the idle
// connection limit will not be exceeded.
// If err != nil, the value of dc is ignored.
// If err == nil, then dc must not equal nil.
// If a connRequest was fulfilled or the *driverConn was placed in the
// freeConn list, then true is returned, otherwise false is returned.
func (db *DB) putConnDBLocked(dc *driverConn, err error) bool {
	if db.closed {
		return false
	}
	if db.maxOpen > 0 && db.numOpen > db.maxOpen {
		return false
	}
	
	if c := len(db.connRequests); c > 0 {
	//如果有等待连接阻塞的请求
		var req chan connRequest
		var reqKey uint64
		for reqKey, req = range db.connRequests {
			break
		}
		delete(db.connRequests, reqKey) // Remove from pending requests.
		if err == nil {
			dc.inUse = true
		}
		req <- connRequest{
			conn: dc,
			err:  err,
		}
		return true
	} else if err == nil && !db.closed {
	//否则将连接放回freeConn
		if db.maxIdleConnsLocked() > len(db.freeConn) {
			db.freeConn = append(db.freeConn, dc)
			db.startCleanerLocked()
			return true
		}
		db.maxIdleClosed++
	}
	return false
}
```
释放的过程：
1. 首先检查下当前归还的连接在使用过程中是否发现已经无效，如果无效则不再放入连接池，然后检查下等待连接的请求数新建连接，类似获取连接时的异常处理，如果连接有效则进入下一步；
2. 检查下当前是否有等待连接阻塞的请求，有的话将当前连接发给最早的那个请求，没有的话则再判断空闲连接数是否达到上限，没有则放入freeConn空闲连接池，达到上限则将连接关闭释放。
3. (只执行一次)启动connectionCleaner协程定时检查feeConn中是否有过期连接，有则剔除。

有个地方需要注意的是，Query、Exec操作用法有些差异：

a.Exec(update、insert、delete等无结果集返回的操作)调用完后会自动释放连接；
b.Query(返回sql.Rows)则不会释放连接，调用完后仍然占有连接，它将连接的所属权转移给了sql.Rows，所以需要手动调用close归还连接，即使不用Rows也得调用rows.Close()，否则可能导致后续使用出错，如下的用法是错误的：
```$xslt
//错误
db.SetMaxOpenConns(1)
db.Query("select * from test")

row,err := db.Query("select * from test") //此操作将一直阻塞

//正确
db.SetMaxOpenConns(1)
r,_ := db.Query("select * from test")
r.Close() //将连接的所属权归还，释放连接
row,err := db.Query("select * from test")
//other op
row.Close()
```