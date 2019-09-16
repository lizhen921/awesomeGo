## 并发和并行

在编程领域，并发（Concurrency）是独立的执行过程 (Process)的组合，而并行（Parallelism)则是计算（可能是相关联的）的同时执行。

并发（Concurrency)是关于同时 应对很多事情(deal with lots of things)

并行（Parallelism)则是同时做许多事情(do lots of things)"

Erlang 之父 Joe Armstrong 的解释，如下图：

![image](https://raw.githubusercontent.com/altairlee/awesomeGo/master/images/goschedule/g-m-p.png)


## 调度器的三个基本对象：
* G (Goroutine)，代表协程，也就是每次代码中使用 go 关键词时候会创建的一个对象
* M (Work Thread)，工作线程
* P (Processor)，代表一个处理器，又称上下文

## G-M-P三者的关系与特点：
1. 每个M必须绑定一个P，线程M创建后会去检查并执行G
2. 每个P都有一个本地G队列，整个调度器还有一个全局的G队列
3. M从队列中提取G并执行
4. P的个数就是GOMAXPROCS（最大256），启动时固定的，一般不修改
5. M的个数和P的个数不一定一样多（会有休眠的M或P不绑定M）（最大10000）
6. P是用一个全局数组（255）来保存的，并且维护着一个全局的P空闲链表

![image](https://raw.githubusercontent.com/altairlee/awesomeGo/master/images/goschedule/g-m-p.png)

## G任务调度

1. Gorutine从入队到执行
```$xslt
    1. 对于新建的协程，它会被加入到本地或者全局队列
    2. 如果有空闲的P（未绑定过M），则创建一个M绑定该P（这里无论是哪个M创建G，都会判断是否有空闲的P，只要有空闲的P就会创建M去绑定）
    3. M会启动一个底层线程，循环执行能找到的G任务
    4. G任务的执行顺序是：本地P的任务队列 --> 其他P的任务队列 --> 全局任务队列
    5. G的切换时间片为10ms，就是说一个协程最多运行10ms，就会被M切换到下一个G（中断，挂起）
    6. 被中断的任务，会放到本地队列队尾，等待再次执行
    7. 中断时，会将寄存器中的栈信息，保存到G对象中，再次执行是，会将自己保存的栈信息复制到寄存器，接着执行。
```

2. P队列
```$xslt
    1. 本地队列没有数据竞争，无需加锁处理，可以提高处理速度
    2. 全局队列是为了保证多个P之间任务平衡。所有M共享全局队列，为保证数据竞争问题，需要加锁处理，相对本地队列处理要慢
```

3. 全局G任务队列和本地队列按照一定的策略进行互换，即`协程交换任务`：
```$xslt
    1.全局与局部：全局G个数/P个数
    2.局部与局部：一次转移一半（Work stealing算法）
```

## 抢占式调度

在每个方法或者函数的入口，增加一段代码，让runtime有机会检查是否需要执行抢占式调度。这种解决方案，对于没有函数调用G，scheduler依然无法抢占




