# Bitmap（位图）

位图不是一种特殊的数据结构，其本质就是string类型，不过redis提供了在string类型上精确到位的操作，当我们按位进行string类型的操作时，把这种类型视为位图。



位图的优势在于其仅仅使用了很少的空间便可以表示很大的信息量，例如需要统计2020年八月有登陆过系统的用户量可以建一个`2020:08:user:sign`的key，每个用户的id视为offset，如果这个用户在八月登录过，就`setbit 2020:08:user:sing  <user id> 1`，当然这种需要用户id为整数（可以映射为offset）。虽然可以使用set数据类型了来做，不过当系统的用户量很大时，set所需要的空间会很大，bitmap占用内存小的优势就体现出来了。



## 基本使用



- `SETBIT key offset value`——设值

- `GETBIT key offset`——取值

- `BITCOUNT key [start end]`——统计统计区间内1的个数。需要注意的是，虽然redis提供了精确到位的设置操作，但是在统计范围时的区间选择上却并不能选择bit位，只能是string字符的位置，也就是说start和end指的是字符的位置。

- `BITPOS key bit [start] [end]`——查询第一个bit出现的offset



### bitfield

前面的`setbit`/`getbit`都是单个位的操作，如果需要一次性操作多个位，需要使用管道或者`bitfield`指令。

`BITFIELD key [GET type offset] [SET type offset value] [INCRBY type offset increment] [OVERFLOW WRAP|SAT|FAIL]`



三个子命令，子命令最多只能处理64个连续的位，如果操作过64位，可以使用使用`bitfield`将其组合一次性执行

- get
- set
- incrby

type指的是有无符号数+取位范围

u代表无符号数，i代表有符号数

u4代表取4位，将结果视为无符号数

i3代表取3位，将结果视为有符号数

e.g. `bitfield w get u4 0 get u3 2 get i4 0 get i3 2`

`overflow`指的是当操作使得key溢出后的处理行为，包括

- wrap——折返（默认）
- sat——饱和截断
- fail——失败不执行