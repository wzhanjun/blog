## 你说说一条更新SQL的执行过程


##### 执行流程

- 客户端发送请求到服务端建立连接
- 服务端先看下查询缓存，对更新某张表的SQL,该表的所有查询缓存都失效
- 接着来到解析器，进行语法分析，一些系统关键字校验，校验语法是否合规
- 然后优化器进行SQL优化（如：怎么选择索引，生成执行计划）
- 执行引擎去存储引擎查询需要更新的数据
- 存储引擎判断当前缓冲池中是否存在需要更新的数据，存在就直接返回，否则去从磁盘加载数据。
- 执行引擎调用存储引擎API去更新数据
- 存储引擎更新数据，同时写入undo_log, redo_log信息
- 执行引擎写binlog,提交事务，流程结束。

##### 流程图



<img src="F:\code\study\github\blog\mysql\433edf1720c7458ca8df7a0b7254bef6.jpeg" alt="433edf1720c7458ca8df7a0b7254bef6" style="zoom:40%;" />





##### redo_log： innodb存储引擎特有  用于事务的原子性和持久性

![4c7a3a71136240b881c92834f7393864](F:\code\study\github\blog\mysql\4c7a3a71136240b881c92834f7393864.jpeg)

在写redo_log的时候先把数据写到redo_log缓冲区，然后异步写入磁盘。

刷盘策略参数：innodb_flush_log_at_trx_commit,   0|1|2，默认的话是1。



0代表提交事务时不会写入磁盘，这样的话性能当然最好，但是在Mysql宕机的情况会丢失上一秒的事务的数据。

1代表提交事务一定会进行一次刷盘，同步当然性能最差，但是也最安全。

2代表写入文件系统的缓存，不进行刷盘。这个选项性能略差于1，Mysql宕机的话对数据没有任何影响，只有在操作系统宕机才会丢失数据，这种情况下默认Mysql每秒会执行一次刷盘。



##### undo_log : 保证了事务的一致性, 保证能够在宕机后恢复事务的数据 

比如insert一条数据，就对应生成delete，update语句则生成相反的更新语句



##### binlog

不同于redo_log是独属于存储引擎独有的东西，binlog则是Mysql本身产生的日志。

redo_log是物理日志，binlog和undo_log都属于逻辑日志。

逻辑日志可以认为就是存储的SQL本身，而物理日志记录的是对页的修改。

另外物理日志可以保证幂等性，而逻辑日志则不一定能，除非本身SQL就是幂等的。

刷盘策略参数：sync_binlog   默认值为0，相当于由文件系统控制（同样如果服务器宕机，binlog丢失）



##### 事务

为了保证写redo_log和binlog的一致性，实际采用了二阶段提交的方式

prepare阶段：根据
innodb_flush_log_at_trx_commit设置的刷盘策略决定是否写入磁盘，标记为prepare状态。

commit阶段：写入binlog日志，事务标记为提交状态。



![4b85e983673344c08ef73cdabb769df7](F:\code\study\github\blog\mysql\4b85e983673344c08ef73cdabb769df7.jpeg)



#### 来源

https://m.toutiao.com/i6997594083476816417

