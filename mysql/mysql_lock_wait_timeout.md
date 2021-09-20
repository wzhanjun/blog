## MySQL Lock 超时


### MySQL造成锁的情况



###  innodb_lock_wait_timeout 和 lock_wait_timeout

- innodb_lock_wait_timeout

    指的是事务等待获取资源等待的最长时间，超过这个时间还未分配到资源则会返回应用失败； 当锁等待超过设置时间的时候，就会报如下的错误；ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction。

    ```
    # 查看
    SHOW VARIABLES LIKE 'innodb_lock_wait_timeout
    # 修改
    set innodb_lock_wait_timeout=100;
    set global innodb_lock_wait_timeout=100;
    # /etc/my.cnf
    innodb_lock_wait_timeout = 50

    ```
- lock_wait_timeout


### 解决方案

``` 
## 应急方法
show full processlist;
kill process_id
## 
select * from innodb_trx;

# 查看innodb_lock_waits
SELECT * FROM innodb_lock_waits;

# innodb_locks 表和 innodb_lock_waits 表结合
SELECT * FROM innodb_locks WHERE lock_trx_id IN (SELECT blocking_trx_id FROM innodb_lock_waits);

# innodb_locks 表 JOIN innodb_lock_waits 表
SELECT innodb_locks.* FROM innodb_locks JOIN innodb_lock_waits ON (innodb_locks.lock_trx_id = innodb_lock_waits.blocking_trx_id);

# 查询 innodb_trx 表
SELECT trx_id, trx_requested_lock_id, trx_mysql_thread_id, trx_query FROM innodb_trx WHERE trx_state = 'LOCK WAIT';

# trx_mysql_thread_id 即kill掉事务线程 ID

```

### 参考

https://ningyu1.github.io/site/post/75-mysql-lock-wait-timeout-exceeded/