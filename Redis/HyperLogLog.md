# HyperLogLog

redis专门提供了一个HyperLogLog数据结构用于解决大数据量下的不精确去重计数，当然使用set来做去重计数也是可以的，而且更加的精准，只不过在数据量非常大的时候占用的内存会很大。HyperLogLog精确度的标准误差是0.81%，占用空间是12K（在数据量很小的情况下，redis进行了优化，并不需要12K。

基于以上特性，HyperLogLog常常被用来计算UV值。

## 基本用法

- `pfadd`——添加元素
- `pfcount`——计算基数值
- `pfmerge`——合并多个HyperLogLog

## 算法原理

TODO



