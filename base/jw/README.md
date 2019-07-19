# Job／Worker工作模式 

基本思路：将多个job放入chan，来缓存需要处理的job；启动多个worker，来消费job chan，并处理。

    MAXJOB：控制job缓存个数
    MAXWORKER：控制最大worker数量
    

TODO：

如果job结果需要返回，怎么处理


参考：

    http://blog.taohuawu.club/article/42#directory0798383148495393810