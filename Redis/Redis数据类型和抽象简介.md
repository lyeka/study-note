#  Redis数据类型和抽象简介

原文地址：[Redis数据类型和抽象简介](https://redis.io/topics/data-types-intro)



Redis实际上是数据结构服务器（*data structures server*）， 支持多种不同类型数据结构，而不是不是简单的键值存储。这意味着，当传统的键值存储服务只能存储字符串与字符串之间的映射时，redis不再限制于存储简单的字符串，有能力处理更加复杂的数据结构。下面是redis支持的数据结构列表，教程后面会分别介绍它们。

- 字符串：二进制安全
- 列表：字符串集合，按元素插入的顺序排列。其存储结构实际上是链表。
- 集合：元素唯一，没有排序的字符串集合
- 有序集合：类似于集合，但是每一个元素都与一个分值（score）相关联。元素总是取分值排序被取出，与集合不同的是有序集合可以按范围取出元素（如取出排前十的元素或者排末尾前十的元素）
- 哈希（Hashes）：键值均为字符串的字典，与ruby、Python中的hashes概念十分接近。
- Bit arrays（bitmaps）：有能力通过特殊命令想处理位数据一样处理字符串，可以单独设置、清除每个位，计算所有设置为1的位，找到第一个设置或者未被只设置的位等等
- HyperLogLogs：这是一个概率数据结构，用于估算集合的基数。别害怕，它比看起来更简单，请参考本教程中相关章节
- Streams：仅追加的字典条目集合，提供了抽象的日志数据类型，在[Introduction to Redis Streams](https://redis.io/topics/streams-intro)有深入介绍

从命令参考中掌握这些数据类型以及如何使用它们来解决问题并不总是那么容易，因此本教程快速讲解redis数据类型以及它们最常用模式。

下面是示例我们将会使用redis-cli工具——一个简单方便的命令行工具，用于发送命令给redis服务器。



## Redis key

redis key是二进制安全的，这意味着你可以使用任何二进制序列作为key，包括字符串如"foo"或是JPEG文件的内容。空字符串同样是有限的key。

一些关于key的规则：

- 非常长的key不是一个明智的做法。例如，一个大小为1024字节的key不仅在内存上不友好，而且在数据集中查找key需要花费更多资源。如果手头的任务中的key非常大， 哈希（例如SHA1算法）后再作为key为更加好，尤其是在内存个带宽的角度来看。
- 非常短的key通常也不是明智的做法。没有理由使用"u1000flw"而不是"user:1000:followers"作为key。后者更加可读并且相对于其存储的来说key占用的内存其实很小。虽然短的key占用的内存更小，但是建议在可读性与内存占用之间取得一个比较好的平衡点。
- 尝试坚持一种模式（作为key的格式）。例如"object-type:id"（<对象>:<id>格式）是一个好的例子， 如"user:100"。常用点或者（英文）破折号作为多字字段的连接符，例如"comment:123:reply.to"或是"comment:123:reply-to"
- key最大允许的大小为512MB

## 字符串

在redis中，字符串是最简单的类型，这是Memcached中唯一支持的数据结构，所以非常对于redis初学者来说使用字符串会非常自然。

redis key 是字符串，因此我们使用字符串时，是将字符串映射到另一个字符串。字符串类型在许多使用场景下非常有用，例如缓存HTMl片段或者页面。

让我们使用redis-cli展示字符串类型的一些使用。

```bash
> set mykey somevalue
OK
> get mykey
"somevalue"
```

如你所见，可以使用SET以及GET命令设置以及取回字符串值。注意SET的key如果事先已经存的话，SET将会覆盖旧值，即使旧值是别的数据类型。

值可以是可以任何类型的字符串（包括二进制数据），例如你可以存储jpeg格式照片作为值。值的大小不可以超过512MB。

SET命令通过额外的参数提供一些有趣的选项。例如，可以要求如何设置的键值已经存在的话，SET操作失败，或者相反，要求Set操作只在键只存在的情况下才成功

```shell
> set mykey newval nx
(nil)
> set mykey newcal xx
OK
```



即使字符串是redis最基本的值，其仍有许多有趣的操作。例如，其中之一便是原子增加。

```shell
> set counter 100
OK
> incr counter
(integer) 101
> incr counter 
(integer) 102
> incrby counter 50
(integer) 152
```

INCR命令解析字符串值为整数，将其加一，并将加一后的值存储。还有一些相似的命令如 INCRBY，DECR、DECRBY。这些命令内在是同一个操作，不过以不同的方式表现。

INCR 命令是原子操作。即使同时有不同的客户端对同一个键值执行INCR操作也不会出现数据竞争。例如，将永远不会出现客户端1和客户端2在同一时间（执行INCR操作）读取值为10，将其增加1到11，以及最终存储为11的情况。如果没有当其他客户端也在这时对该执行INCR操作的话，该值最终的值将会为12。

字符串有大量的相关命令。例如，GETSET命令在更新值的同时返回旧值。例子略

批量设置或者取回多个键的值操作在减少时延方面很有效。因此衍生出MSET和MEGT命令：

```shell
> mset a 10 b 20 c 30
OK
> mget a b c
1)  "10"
2) "20"
3) "30"
```

MGET 会以数组的方式返回值



### Altering and querying the key space

有一些命令没有限定于某种数据类型。但在检索键的空间方面有用，所以这些命令适用于任何类型

如EXISTS命令将根据键的存在与否返回1或0，DEL命令删除一个键与其值，无论值是什么。

```shell
> set mykey hello
OK
> exists mykey
(integer) 1
> del mykey
(integer) 0
> exists mykey
(integer) 0
```

DEL根据删除的键存在与否返回1或0。

有许多与键空间相关的命令，但是上述两个以及TYPE命令是最重要的。TYPE命令返回值的类型。

```shell
> set mykey x
OK
> type mykey
string
> del mykey
(integer) 1
> type mykey
none
```



### Redis 过期机制： 带有限生存时间的键

在深入其他复杂的数据结构之前，我们需要讨论下另外一个与类型无关的特性——redis过期机制。基本上你就可以为键设置一个超时时间，用于限定键的生存时间，当生存时间达到之后，键将会自动销毁，就如同用户调用了DEL命令一样。

关于redis过期机制的一些快速概览

- 时间精度可以是秒或者毫秒
- 但是， 到期时间的分辨率始终为1毫秒
- 关于过期的信息被复制并保留在磁盘上，即使redis关机了生存时间也在继续消耗（这意味着redis保存的是键的实际到期时间点）

设置过期时间很简单

```shell
> set key some-value
OK
> expire key 5
(integer) 1
> get key (immediately)
"some value"
> get key (after some time)
(nil)
```

键在两次get调用之间消失了，因为两次调用的时间差超过了5秒。在上诉例子中，我们使用EXPIRE设置过期时间（同样可以为已经有过期时间的键设置不同的过期机制，像PERSIST命令将会移除键的过期时间），然而我们同样可以使用别的命令设置过期时间，例如在GET中使用额外的参数

```shell
> set key 100 ex 10
OK
> ttl key
(integer) 9
```

上述例子将键的值设置为100，生存时间为10秒。后面的TTL命令用于查看键的剩余生成时间。

如果想设置或者查看毫秒精度的生存时间，使用PEXPIRE和PTTL命令，以及SET的其余选项。







疑问

- redis expire 可靠吗（实现原理）以及 expire 策略？

ref

- [二进制安全(binary safe)是什么意思？](https://www.zhihu.com/question/28705562)

