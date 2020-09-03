# Scan

redis `keys`指令可以用于查找符合满足给定正则字符串规则的key，但是存在一些缺点

- 没有limit， offset参数，输出结果可能非常多
- 复杂度为O(n)，key指令太多的话会导致redis服务卡顿，阻塞其他指令执行（redis单线程）

为此redis提供了`scan`指令，提供了通过游标迭代key



## 基础使用

- `scan <cursor> match <rule> count <limit>`

返回值包括

- 游标值
    - 用于下次`scan`使用，但游标值为0时代表没有更多数据
- 符合规则的key列表，其数量不一定等于limit值



## 原理

![img](https://user-gold-cdn.xitu.io/2018/7/5/164695b9f06c757e?imageView2/0/w/1280/h/960/format/webp/ignore-error/1)

`scan`的迭代顺序不是递增的，这是考虑到可保存key的字典（如上图）存在扩缩容情况，如果是递增的，会存在遍历重复和遗漏的情况。`scan`采取的是高位进位加法，如下图示



![img](https://user-gold-cdn.xitu.io/2018/7/5/16469760d12e0cbd?imageslim)

字典扩缩容后，将会进行rehash，保存key的槽会变动。在以2^n方式扩缩容时，槽位的话简单来说就会发生如下变动

> 假设开始槽位的二进制数是 xxx，那么该槽位中的元素将被 rehash 到 0xxx 和 1xxx(xxx+8) 中。 如果字典长度由 16 位扩容到 32 位，那么对于二进制槽位 xxxx 中的元素将被 rehash 到 0xxxx 和 1xxxx(xxxx+16) 中。

![img](https://user-gold-cdn.xitu.io/2018/7/5/164699dae277cc19?imageView2/0/w/1280/h/960/format/webp/ignore-error/1)

在这种rehash下，高位进位加法的迭代去避免了重复和遗漏的问题。



### 渐进式rehash

当字典元素过多时，rehash会导致占用时间过长，所以redis采用的是渐进式rehash，也即是

> 它会同时保留旧数组和新数组，然后在定时任务中以及后续对 hash 的指令操作中渐渐地将旧数组中挂接的元素迁移到新数组上。这意味着要操作处于 rehash 中的字典，需要同时访问新旧两个数组结构。如果在旧数组下面找不到元素，还需要去新数组下面去寻找。

`scan`同样也要考虑这个问题



## 使用场景

扫描redis占用空间太大的key

解决方法

1. 脚本使用`scan`，计算key对应的size或者len， 将排名靠前的展示出来
2. 使用官方指令——`redis-cli  –-bigkeys`



ref

- https://juejin.im/book/6844733724618129422/section/6844733724710404110