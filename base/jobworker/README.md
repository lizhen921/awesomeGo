# Job／Worker工作模式

基本思路：job是一个chan，来缓存需要处理的job；启动多个worker，来消费job chan，并处理。

