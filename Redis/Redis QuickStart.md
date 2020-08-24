# Redis QuickStart



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



## List

- `BLPOP`——取出列表第一个元素，没有元素的话会阻塞
- `BRPOP`——取出列表最后一个元素，没有元素的话阻塞
- `BRPOPLPUSH`——去除列表最后一个元素，插入到另外一个列表，没有元素的话阻塞
- `LINDEX`——通过索引获取列表的元素
- `LINSERT`——在列表元素的前面或者后面插入元素
- `LLEN`——获取元素长度
- `LPOP`——移除列表的第一个元素
- `LPUSH`——插入列表头部
- `LPUSHX`——插入到已存在列表头部
- `LRANGRE`——获取指定范围内的列表元素
- `LREM`——移除列表元素
- `LSET`——通过索引设置列表元素值
- `LTRIM`——对列表进行修剪，保留指定区间内的元素
- `RPOP`——移除列表最后一个元素
- `RPOPLPUSH`——移除列表最后一个元素，插入到另外一个元素
- `RPUSH`——插入元素到列表末尾
- `RPUSHX`——插入元素到已存在的列表末尾



## Set

无序集合，集合成员唯一。通过哈希表来实现，所以添加，删除，查找复杂度都是O(1)

- `SADD`——添加成员
- `SCARD`——获取集合成员数
- `SDIFF`——返回第一个集合与其他集合的差异
- `SDIFFSTORE`——保存第一个集合与其他集合的差异为新集合
- `SINTER`——计算集合交集
- `SINTERSTORE`——保存集合交集为新集合
- `SISMENBER`——判断成员是否在集合中
- `SMEMBERS`——返回集合中成员
- `SMOVE`——将集合成员移动到另外的集合
- `SPOP`——随机移除并返回一个成员
- `RANDOMMEMBER`——返回集合中指定数量的成员
- `SREM`——删除指定成员
- `SUNION`——计算集合并集
- `SUNIONSTORE`——保存集合并集并保存为新集合
- `SSCAN`——迭代集合中的元素



## Sort Set

特性与集合差不多，每个成员都会关联一个double类型的分数，分数可以重复

下面所说的索引可以认为是排名（从小到大排序）

- `ZADD`——添加成员
- `ZCARD`——获取集合成员数
- `ZCOUNT`——计算指定区间分数的成员数
- `ZINCRBY`——对成员分数加上增量
- `ZINTERSTORE`——计算交集，保存在新集合中
- `ZLEXCOUNT`——计算指定字典区间内成员数量
- `ZRANGE`——通过索引返回区间内成员
- `ZREVRANGE`——通过索引区间返回成员（逆序）
- `ZRANGEBYLEX`——通过字典区间返回有序集合成员
- `ZRANGEBYSCORE`——通过分数区间返回成员
- `ZRANK`——返回成员的索引
- `ZREM`——移除成员
- `ZREMRANGEBYLEX`——通过字典区间移除成员
- `ZREMRANGEBYRANK`——通过索引区间移除成员
- `ZREMRANGEBYSCORE`——通过分数区间移除成员
- `ZREVRANK`——返回成员索引（逆序）
- `ZSCORE`——返回成员分数
- `ZUNIONSTORE`——计算并集，保存在新集合
- `ZSCAN`——迭代有序集合元素



## HyperLogLog

用于计算基数（集合中包含元素的个数）

优点——在输入元素数量非常大时，计算基数所需要的空间总是固定的，并且是很小的。

在Redis里面，每个HyperLogLog键只需要话费12KB内存，就可以计算接近2^64个不同元素的基数。



- `PFADD`——添加元素

- `PFCOUNT`——返回基数估算值

- `PFMERGE`——合并多个HyperLogLog为一个HyperLogLog

    

## 发布订阅

发布者（pub）发送消息，订阅者（sub）接收消息

- `SUBSCRIBE`——订阅一个或者多个频道
- `UNSUBSCRIBE`——退订给定的频道

- `PSUBSCRIBE`——订阅一个或者多个给定模式的频道

- `PUNSUBSCRIBE`——退订所有给定模式的频道

- `PUBLISH`——将消息发送道指定的频道

- `PUBSUB`——查看订阅与发布系统状态

    



## 事务

Redis单个操作是原子性的，提供了“不完整”的事务支持

执行过程：

1. 使用`MULTI`开始事务，而后的命令会被放进队列缓存
2. 使用`EXEC`结束缓存命令，命令按顺序执行，其中命令执行的失败不会停止剩余命令的执行，前面已经执行的命令也不会回退



- `MULTI`——标记一个事务的开始
- `EXEC`——执行事务所有命令
- `DISCARD`——取消事务
- `WATCH`——监控一个或者多个key，如果在事务执行前这个key被其他命令锁改动，那么事务被打断（不再执行，返回错误）
- `UNWATCH`——取消watch对所有key的监控



ref

- [redis位操作](https://www.cnblogs.com/xuwenjin/p/8885376.html)
- [有序集合中的LEX指什么](https://www.twle.cn/l/yufei/redis/redis-basic-sorted-sets-zrangebylex.html)
- [redis事务](https://redisbook.readthedocs.io/en/latest/feature/transaction.html)

