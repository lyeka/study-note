

## 字符串

- `SET`——设置指定的key值
- `GET`——获取指定的key值
- `GETRANGE`——返回子字符串
- `GETSET`——设置新值，返回旧值
- `GETBIT`——获取字符串指定偏移量上的位（bit）
- `MGET`——批量获取值
- `SETBIT`——设置指定位置的位
- `SETEX`——设置指定值，并指定过期时间（秒）
- `SETNX`——只有key不存在时才设置值
- `SETRANGE`——设置子字符串
- `STRLEN`——返回值长度
- `MSET`——批量设置
- `MSETNX`——批量设置值，仅当key不存在时
- `PSETEX`——设置值，指定过期时间（毫秒）
- `INCR`——加一
- `INCRBY`——加上指定整数
- `INCRBYFLAOT`——加上指定浮点值
- `DECR`——减一
- `DECRBY`——减去指定值
- `APPEND`——字符串末尾添加值



## Hash

string类型的field（字段）和value（值）映射表，适合用于存储对象。

- `HDEL`——删除一个或多个哈希字段
- `HEXISTS` ——查看指定字段是否存在
- `HGET`——获取字段值
- `GHETALL`——获取key所有字段和值
- `HINCRBY`——为指定字段加上指定整数数值
- `HINCRBYFLOAT`——为指定字典加上指定浮点值
- `HKEYS`——获取哈希所有字段
- `HLENS`——获取字段数量
- `HMGET`——批量获取字段值
- `HMSET`——批量设置字段值
- `HSET`——设置字段值
- `HSETNX`——字段不存在是才设置字段值
- `HVALS`——获取哈希所有值
- `HSCAN`——迭代哈希中的键值对



ref

- [redis位操作](https://www.cnblogs.com/xuwenjin/p/8885376.html)

