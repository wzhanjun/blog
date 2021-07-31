## Lua 在redis中的运用



#### EVAL 命令

- `EVAL` 命令的关键字。

- `luascript` Lua 脚本。

- `numkeys` 指定的Lua脚本需要处理键的数量，其实就是 `key`数组的长度。

- `key` 传递给Lua脚本零到多个键，空格隔开，在Lua 脚本中通过 `KEYS[INDEX]`来获取对应的值，其中`1 <= INDEX <= numkeys`。

- `arg`是传递给脚本的零到多个附加参数，空格隔开，在Lua脚本中通过`ARGV[INDEX]`来获取对应的值，其中`1 <= INDEX <= numkeys`。

  

```
EVAL luascript numkeys key [key ...] arg [arg ...]

# get
EVAL "return redis.call('GET',KEYS[1])" 1 hello  
# set
EVAL "return redis.call('SET', KEYS[1], 'abc')" 1 "test"  

```

**注意**：

1. **Redis和Lua互相传递数据时存在类型转换，注意精度丢失， 转换为字符串则是安全的**

   ```
   EVAL "return 3.14" 0  ====>  3
   EVAL "return tostring(3.14)" 0    "3.14"
   ```

2. **执行Lua脚本时其它客户端发送的命令将被阻塞**





#### 参考

[https://segmentfault.com/a/1190000037518418](https://segmentfault.com/a/1190000037518418)