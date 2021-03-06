# MongoDB

[TOC]



## 安装

版本 v4.2

```shell
wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | sudo apt-key add -
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.2.list
sudo apt-get update
sudo apt-get install -y mongodb-org
```



## CRUD

### 插入

#### 插入单条文档

`db.collection.insertOne()` 

```shell
> db.inventory.insertOne(
   { item: "canvas", qty: 100, tags: ["cotton"], size: { h: 28, w: 35.5, uom: "cm" } }
)

# output
{
	"acknowledged" : true,
	"insertedId" : ObjectId("5eaa9eb532b6e77e797bc19c")
}
```

#### 插入多条文档

`db.collection.insertMany()`

```shell
> db.inventory.insertMany([
   { item: "journal", qty: 25, tags: ["blank", "red"], size: { h: 14, w: 21, uom: "cm" } },
   { item: "mat", qty: 85, tags: ["gray"], size: { h: 27.9, w: 35.5, uom: "cm" } },
   { item: "mousepad", qty: 25, tags: ["gel", "blue"], size: { h: 19, w: 22.85, uom: "cm" } }
])

# output
{
	"acknowledged" : true,
	"insertedIds" : [
		ObjectId("5eaaa35af6405bb84d1953a0"),
		ObjectId("5eaaa35af6405bb84d1953a1"),
		ObjectId("5eaaa35af6405bb84d1953a2")
	]
}
```



#### 插入文档

`db.collection.insert()`

插入单或者多个文档，取决与传入的是单个文档还是文档数组



#### 插入行为

- 集合不存在的话，自动创建集合
- mongodb没条文档都需要一个唯一的`_id`字段值作为主键，插入文档不存在`_id`字段的话，会自动生成插入
- 只支持单个文档的原子性写操作（即批量操作不支持整体操作的原子性）
- Write Acknowledgement #todo



### 查询

先用下列文档填好数据以作示例

```shell
db.inventory.insertMany([
   { item: "journal", qty: 25, size: { h: 14, w: 21, uom: "cm" }, status: "A" },
   { item: "notebook", qty: 50, size: { h: 8.5, w: 11, uom: "in" }, status: "A" },
   { item: "paper", qty: 100, size: { h: 8.5, w: 11, uom: "in" }, status: "D" },
   { item: "planner", qty: 75, size: { h: 22.85, w: 30, uom: "cm" }, status: "D" },
   { item: "postcard", qty: 45, size: { h: 10, w: 15.25, uom: "cm" }, status: "A" }
]);
```

#### 查询所有文档

`db.inventory.find( {} )`

```shell
> db.inventory.find( {} )

# output
{ "_id" : ObjectId("5eaaa35af6405bb84d1953a0"), "item" : "journal", "qty" : 25, "tags" : [ "blank", "red" ], "size" : { "h" : 14, "w" : 21, "uom" : "cm" } }
{ "_id" : ObjectId("5eaaa35af6405bb84d1953a1"), "item" : "mat", "qty" : 85, "tags" : [ "gray" ], "size" : { "h" : 27.9, "w" : 35.5, "uom" : "cm" } }
{ "_id" : ObjectId("5eaaa35af6405bb84d1953a2"), "item" : "mousepad", "qty" : 25, "tags" : [ "gel", "blue" ], "size" : { "h" : 19, "w" : 22.85, "uom" : "cm" } }
{ "_id" : ObjectId("5eaaa8a7f6405bb84d1953a3"), "item" : "journal", "qty" : 25, "size" : { "h" : 14, "w" : 21, "uom" : "cm" }, "status" : "A" }
{ "_id" : ObjectId("5eaaa8a7f6405bb84d1953a4"), "item" : "notebook", "qty" : 50, "size" : { "h" : 8.5, "w" : 11, "uom" : "in" }, "status" : "A" }
{ "_id" : ObjectId("5eaaa8a7f6405bb84d1953a5"), "item" : "paper", "qty" : 100, "size" : { "h" : 8.5, "w" : 11, "uom" : "in" }, "status" : "D" }
{ "_id" : ObjectId("5eaaa8a7f6405bb84d1953a6"), "item" : "planner", "qty" : 75, "size" : { "h" : 22.85, "w" : 30, "uom" : "cm" }, "status" : "D" }
{ "_id" : ObjectId("5eaaa8a7f6405bb84d1953a7"), "item" : "postcard", "qty" : 45, "size" : { "h" : 10, "w" : 15.25, "uom" : "cm" }, "status" : "A" }
```

#### 按字段查询

`db.inventory.find({ <field1>: <value1>, ... })`

```shell
> db.inventory.find( { status: "D" } )
# output
{ "_id" : ObjectId("5eaaa8a7f6405bb84d1953a5"), "item" : "paper", "qty" : 100, "size" : { "h" : 8.5, "w" : 11, "uom" : "in" }, "status" : "D" }
{ "_id" : ObjectId("5eaaa8a7f6405bb84d1953a6"), "item" : "planner", "qty" : 75, "size" : { "h" : 22.85, "w" : 30, "uom" : "cm" }, "status" : "D" }
```



#### 使用查询运算符

`db.inventory.find({ <field1>: { <operator1>: <value1> }, ... })`

```shell
> db.inventory.find( { status: { $in: [ "A", "D" ] } } )
# output
{ "_id" : ObjectId("5eaaa8a7f6405bb84d1953a3"), "item" : "journal", "qty" : 25, "size" : { "h" : 14, "w" : 21, "uom" : "cm" }, "status" : "A" }
{ "_id" : ObjectId("5eaaa8a7f6405bb84d1953a4"), "item" : "notebook", "qty" : 50, "size" : { "h" : 8.5, "w" : 11, "uom" : "in" }, "status" : "A" }
{ "_id" : ObjectId("5eaaa8a7f6405bb84d1953a5"), "item" : "paper", "qty" : 100, "size" : { "h" : 8.5, "w" : 11, "uom" : "in" }, "status" : "D" }
{ "_id" : ObjectId("5eaaa8a7f6405bb84d1953a6"), "item" : "planner", "qty" : 75, "size" : { "h" : 22.85, "w" : 30, "uom" : "cm" }, "status" : "D" }
{ "_id" : ObjectId("5eaaa8a7f6405bb84d1953a7"), "item" : "postcard", "qty" : 45, "size" : { "h" : 10, "w" : 15.25, "uom" : "cm" }, "status" : "A" }

```

```shell
# 查询运算符示例
#  in 
db.inventory.find( { status: { $in: [ "A", "D" ] } } )
# and 直接多个键值对即可，无and关键字
db.inventory.find( { status: "A", qty: { $lt: 30 } } ) # 等同于sql中的  SELECT * FROM inventory WHERE status = "A" AND qty < 30
# or
db.inventory.find( { $or: [ { status: "A" }, { qty: { $lt: 30 } } ] } )
# and or 混用
db.inventory.find( {status: "A",$or: [ { qty: { $lt: 30 } }, { item: /^p/ } ]} ) # 等同于 SELECT * FROM inventory WHERE status = "A" AND ( qty < 30 OR item LIKE "p%")

```

[查询运算符参考](https://docs.mongodb.com/manual/reference/operator/query/)



#### 单条查询

`db.collection.findOne`

等同于`db.collection.find()`限制limit=1



#### 嵌套文档字段查询

嵌套文档字段之间使用`.`号连接，查询方法和普通查询没差别

```shell
db.inventory.find( { "size.h": { $lt: 15 } } )
```

上例表示查询size字段中h字段值小于15的记录



#### 在数组字段上查询

首先使用下列数据填充

```shell
db.inventory.insertMany([
   { item: "journal", qty: 25, tags: ["blank", "red"], dim_cm: [ 14, 21 ] },
   { item: "notebook", qty: 50, tags: ["red", "blank"], dim_cm: [ 14, 21 ] },
   { item: "paper", qty: 100, tags: ["red", "blank", "plain"], dim_cm: [ 14, 21 ] },
   { item: "planner", qty: 75, tags: ["blank", "red"], dim_cm: [ 22.85, 30 ] },
   { item: "postcard", qty: 45, tags: ["blue"], dim_cm: [ 10, 15.25 ] }
]);
```



##### 通过数组筛选

使用数组匹配数组，两者的元素值以及顺序都需要完全一致

```shell
db.inventory.find( { tags: ["red", "blank"] } )
```

如果不需要顺序一致，只需要值一致，使用`$all`操作符

```shell
db.inventory.find( { tags: { $all: ["red", "blank"] } } )
```

##### 通过数组元素筛选

查询数组值中包含该值

```shell
db.inventory.find( { tags: "red" } )
```

使用查询操作符

```shell
db.inventory.find( { dim_cm: { $gt: 25 } } )
```

指定多个条件

**数组元素满足条件其中之一即可**

```shell
db.inventory.find( { dim_cm: { $gt: 15, $lt: 20 } } )
```

**使用`$elemMatch`操作符，数组元素需要有一个元素满足所有查询条件**

```shell
db.inventory.find( { dim_cm: { $elemMatch: { $gt: 22, $lt: 30 } } } )
```

##### 通过数组索引元素筛选

使用点标识法，字段与位置索引之间使用`.`连接（即`field1.1` 代表数组field1字段的第二个元素）表示数组元素，查询方法与普通查询没差

##### 通过数值长度筛选

使用`$size`操作符标识数组长度限制

```shell
db.inventory.find( { "tags": { $size: 3 } } )
```



#### 在嵌套文档数组字段上查询

填充实例数据

```shell
db.inventory.insertMany( [
   { item: "journal", instock: [ { warehouse: "A", qty: 5 }, { warehouse: "C", qty: 15 } ] },
   { item: "notebook", instock: [ { warehouse: "C", qty: 5 } ] },
   { item: "paper", instock: [ { warehouse: "A", qty: 60 }, { warehouse: "B", qty: 15 } ] },
   { item: "planner", instock: [ { warehouse: "A", qty: 40 }, { warehouse: "B", qty: 5 } ] },
   { item: "postcard", instock: [ { warehouse: "B", qty: 15 }, { warehouse: "C", qty: 35 } ] }
]);
```

##### 查询嵌套在数组中的文档

数组元素中至少有一个文档与查询文档完全一致（包括字段顺序）

```shell
db.inventory.find( { "instock": { warehouse: "A", qty: 5 } } )
```

##### 在文档数组中的字段上指定查询条件

一、 在嵌套文档数组中的字段上指定查询条件

数组元素中至少有一个文档满足其字段符合条件

注意，这里的文档字段需要与数组字段用`.`连接，并且使用引号包围

```shell
db.inventory.find( { 'instock.qty': { $lte: 20 } } )
```

二、 使用数组索引在嵌套文档中查询字段

嵌套文档数组对应索引位置的文档满足筛选条件

```shell
db.inventory.find( { 'instock.0.qty': { $lte: 20 } } )
```

##### 为文档数组指定多个条件

一、单个嵌套文档在嵌套字段上满足多个查询条件

数组元素中至少有一个元素同时符合筛选条件

```shell
db.inventory.find( { "instock": { $elemMatch: { qty: 5, warehouse: "A" } } } )
```

```shell
db.inventory.find( { "instock": { $elemMatch: { qty: { $gt: 10, $lte: 20 } } } } )
```

二、 元素组合满足筛选条件

数组至少存在文档对应字段字段满足条件之一（不限定是同一份文档，数组中有任意文档任意字段满足筛选条件即可）

```shell
db.inventory.find( { "instock.qty": { $gt: 10,  $lte: 20 } } )
```

```shell
db.inventory.find( { "instock.qty": 5, "instock.warehouse": "A" } )
```

#### 限定查询返回字段

填充实例数据

```shell
db.inventory.insertMany( [
  { item: "journal", status: "A", size: { h: 14, w: 21, uom: "cm" }, instock: [ { warehouse: "A", qty: 5 } ] },
  { item: "notebook", status: "A",  size: { h: 8.5, w: 11, uom: "in" }, instock: [ { warehouse: "C", qty: 5 } ] },
  { item: "paper", status: "D", size: { h: 8.5, w: 11, uom: "in" }, instock: [ { warehouse: "A", qty: 60 } ] },
  { item: "planner", status: "D", size: { h: 22.85, w: 30, uom: "cm" }, instock: [ { warehouse: "A", qty: 40 } ] },
  { item: "postcard", status: "A", size: { h: 10, w: 15.25, uom: "cm" }, instock: [ { warehouse: "B", qty: 15 }, { warehouse: "C", qty: 35 } ] }
]);
```

##### 返回所有字段

默认行为

##### 返回指定字段

使用投影文档（[projection](https://docs.mongodb.com/manual/reference/glossary/#term-projection) document）限制返回的字段

```shell
db.inventory.find( { status: "A" }, { item: 1, status: 1 } )
```

##### 不返回指定字段

```shell
db.inventory.find( { status: "A" }, { status: 0, instock: 0 } )
```

##### 不返回`_id`字段

上述限定字段还默认返回`_id`字段

显示指定`_id`为0改变这一行为

```shell
db.inventory.find( { status: "A" }, { item: 1, status: 1, _id: 0 } )
```

```shell
db.inventory.find( { status: "A" }, { status: 0, instock: 0, _id: 0 } )
```

注意，`_id`限定为1是会报错的



##### 限定嵌套文档的字段显示

使用点标识表示嵌套文档的字段

```shell
db.inventory.find(
   { status: "A" },
   { "size.uom": 0 }
)
```



##### 限定嵌套文档数组字段显示

使用点标识表示嵌套文档的字段

```shell
db.inventory.find( { status: "A" }, { item: 1, status: 1, "instock.qty": 1 } )
```

#####  限定数组元素中特定元素

`$slice`

返回数组

```shell
db.inventory.find( { status: "A" }, { item: 1, status: 1, instock: { $slice: -1 } } )
```



`$`

[传送门](https://docs.mongodb.com/manual/reference/operator/projection/positional/#proj._S_)

todo



`$elemMatch`

[传送门](https://docs.mongodb.com/manual/reference/operator/projection/elemMatch/#proj._S_elemMatch)

todo



#### 查询Null或者不存在值

填充实例数据

```shell
db.inventory.insertMany([
   { _id: 1, item: null },
   { _id: 2 }
])
```

##### 查询Null或者不存在值

下列查询会返回item字段为null 或者 不存在item字段的文档

```shell
db.inventory.find( { item: null } )
```

##### 限定只查询Null值

下列查询加上了类型检测 `{$type: 10}` 即为null值

```shell
db.inventory.find( { item : { $type: 10 } } )
```

[BSON类型传送门](https://docs.mongodb.com/manual/reference/bson-types/)



##### 查询字段存在与否

```shell
db.inventory.find( { item : { $exists: false } } )
```



#### 在mongodb shell中迭代游标

[传送门](https://docs.mongodb.com/manual/tutorial/iterate-a-cursor/)

只说了mongo shell中的行为，不清楚在编程语言中的表现，这个只给传送门，后续了解再加



### 更新

填充实例数据

```shell
db.inventory.insertMany( [
   { item: "canvas", qty: 100, size: { h: 28, w: 35.5, uom: "cm" }, status: "A" },
   { item: "journal", qty: 25, size: { h: 14, w: 21, uom: "cm" }, status: "A" },
   { item: "mat", qty: 85, size: { h: 27.9, w: 35.5, uom: "cm" }, status: "A" },
   { item: "mousepad", qty: 25, size: { h: 19, w: 22.85, uom: "cm" }, status: "P" },
   { item: "notebook", qty: 50, size: { h: 8.5, w: 11, uom: "in" }, status: "P" },
   { item: "paper", qty: 100, size: { h: 8.5, w: 11, uom: "in" }, status: "D" },
   { item: "planner", qty: 75, size: { h: 22.85, w: 30, uom: "cm" }, status: "D" },
   { item: "postcard", qty: 45, size: { h: 10, w: 15.25, uom: "cm" }, status: "A" },
   { item: "sketchbook", qty: 80, size: { h: 14, w: 21, uom: "cm" }, status: "A" },
   { item: "sketch pad", qty: 95, size: { h: 22.85, w: 30.5, uom: "cm" }, status: "A" }
] );
```

#### 更新单条文档

` db.collection.updateOne()`更新第一条符合筛选条件的文档

下列操作更新第一条item字段为paper的文档，使用$set操作符更新对应字段值， $currentDate操作符将lastModified字段更为当前日期，如果该字段不存在则创建

```shell
db.inventory.updateOne(
   { item: "paper" },
   {
     $set: { "size.uom": "cm", status: "P" },
     $currentDate: { lastModified: true }
   }
)
```



#### 更新多条文档

` db.collection.updateMany()` 更新所有符合筛选条件的文档

```shell
db.inventory.updateMany(
   { "qty": { $lt: 50 } },
   {
     $set: { "size.uom": "in", status: "P" },
     $currentDate: { lastModified: true }
   }
)
```



#### 替换文档

`db.collection.replaceOne()`替换文档内容，除了_id字段

如果替换文档中包含`_id`字段,其必须与筛选文档的`_id`字段保持一致

```shell
db.inventory.replaceOne(
   { item: "paper" },
   { item: "paper", instock: [ { warehouse: "A", qty: 60 }, { warehouse: "B", qty: 40 } ] }
)
```





#### 使用聚合管道更新

todo



#### 更新行为

- 保证单条文档操作的原子性
- 不可以更新`_id`字段/不可以替换`_id`不一致的文档
- 保留写操作的字段顺序，除了`_id`字段以及使用字段名称更改操作
- 使用`upsert: true`在没找到对应筛选文档时插入一条新文档
- 指定写入操作确认级别 todo



### 删除

填充示例数据

```shell
db.inventory.insertMany( [
   { item: "journal", qty: 25, size: { h: 14, w: 21, uom: "cm" }, status: "A" },
   { item: "notebook", qty: 50, size: { h: 8.5, w: 11, uom: "in" }, status: "P" },
   { item: "paper", qty: 100, size: { h: 8.5, w: 11, uom: "in" }, status: "D" },
   { item: "planner", qty: 75, size: { h: 22.85, w: 30, uom: "cm" }, status: "D" },
   { item: "postcard", qty: 45, size: { h: 10, w: 15.25, uom: "cm" }, status: "A" },
] );
```

#### 删除多条文档

删除所有文档

```shell
db.inventory.deleteMany({})
```



删除所有符合筛选条件的文档

```shell
db.inventory.deleteMany({ status : "A" })
```



#### 删除单条文档

```shell
db.inventory.deleteOne( { status: "D" } )
```



#### 删除行为

- 删除文档不会删除索引
- 保证单条文档操作原子性
- 确认级别可指定





## 索引

使用索引可以提高查询的效率。

索引是以易于遍历的形式用来存储数据集的一小部分数据的数据结构，其存储一个或多个特定字段的值，按该字段的值排序。

索引排序支持高效等值匹配以及范围查询操作，另外mongodb可以使用索引的排序返回排序好的结果。

mongodb的索引为B-tree数据结构。



**默认的`_id`索引**

mongo会自定为文档的`_id`字段加上索引，这一行为不可改变，且该索引还是唯一索引。



**创建索引**

`db.collection.createIndex( <key and index type specification>, <options> )`

```shell
db.collection.createIndex( { name: -1 } )
```

对于已经存在的索引，创建多次保证幂等性，即使使用别的名称，也无法创建，保留原名称，故索引改名只可以先删除再创建。



**索引名称**

索引默认名称以字段名+排序方向和下划线串联如`{ item : 1, quantity: -1 }`的默认索引名为``item_1_quantity_-1``

可以显示声明来指定索引名称

```shell
db.products.createIndex(
  { item: 1, quantity: -1 } ,
  { name: "query for inventory" }
)
```



**查看集合索引信息**

`db.collection.getIndexes()`



**索引类型**

- 单字段索引
- 复合索引
- 多key索引
- 地理空间索引
- 文本索引
- 哈希索引



**索引属性分类**

**唯一索引**

字段值唯一

**部分索引**

只对符合指定过滤条件的的文档建立索引

**稀疏索引**

跳过没有索引字段的文档

**TTL索引**

具有生命周期，到期会自动删除文档的索引



**分析索引**

[todo](https://docs.mongodb.com/manual/tutorial/analyze-query-plan/)



**索引和排序规则**

[todo]



**覆盖查询**

如果查询条件和返回文档字段仅包含索引字段，mongo可以无需扫描原文档以及在内存中处理二直接返回结果，极大提高查询效率



**索引交集**

对于复合查询条件的查询，如果查询条件分别处于不同的索引之中，mongo可以使用索引交集来满足查询



**索引限制**

[todo]



**其他注意事项**

[todo]

### 单字段索引

![单字段索引存储图示](https://docs.mongodb.com/manual/_images/index-ascending.bakedsvg.svg)

对于单字段来说，索引的排序方向并不重要，因为mongo支持两端遍历索引



### 复合索引

![](https://docs.mongodb.com/manual/_images/index-compound-key.bakedsvg.svg)



对于复合字段，字段的顺序具有重要意义。mongo会按照字段的顺序作优先级来排序索引。

对于复合索引和排序操作，索引键的排序方向可以确定索引是否可以支持排序操作[todo](https://docs.mongodb.com/manual/core/index-compound/#index-ascending-and-descending)



### 多key索引

![](https://docs.mongodb.com/manual/_images/index-multikey.bakedsvg.svg)

如果索引字段存储的是数组，mongo会确定字段为多key索引，为数组的每个元素创建单独的索引条目



### 地理空间索引

[todo]



###  文本索引

支持在文档中搜索字符串内容

[todo]



### 哈希索引

[todo]

## 安全配置



mongodb默认是关闭访问控制的，也就是说直接使用`mongo`就可以登录mongodb, 而不需要像例如mysql那样还需要用户密码登录等。

但是这样不安全，所以一般生产环境需要开启访问控制。

开启访问控制需要在配置文件（ubuntu下是在/etc/mongod.conf文件）加上`authorization: enabled`， 需要注意的是，这个`authorization: enabled`是在 `security`配置下的， 所以还需要把`security`配置给取消注释。

但是在开启访问控制之前得先创建用户，不然开启了访问控制你没得用户去登录。。所以第一步我们应该先不开启访问控制，登录mongodb创建用户

### 创建用户

一般需要创建个超级管理员的用户，地位类似于例如mysql中的root用户

```shell
use admin
db.createUser(
  {
    user: "<user>",
    pwd:  "<password>"
    roles: [ { role: "userAdminAnyDatabase", db: "admin" }, "readWriteAnyDatabase" ]
  }
)
```

mongodb中的用户验证是与db相关联的，在登录选项中要指定这个db，mongodb会使用这个db验证用户/密码，而这个db就只创建用户时使用的db，例如上文就是admin。

尽管这个db与用户验证相关联，但是创建的用户可以在其他db扮演角色，也就是说这个验证db仅仅用于验证用户而已，不限制创建的用户在任何db的权限，限制用户在db中的角色是在document中的roles字段定义的，例如在上文中授予了用户在admin这个db中的角色为`userAdminAnyDatabase`，在这里指定为其他不是admin的db也是没问题的。



创建普通用户

mysql中一般不会使用root来作为业务上连接数据库的用户，而是创建一个普通用户，在给予用户访问哪些库的权利。

mongodb同样可以这样操作。

创建用户的方法和上面创建超级管理员的方法是一样的，不同的是roles角色而已。不过需要注意的是记得使用use选择到你想要验证用户的db，再创建用户。



至于角色类型参考官方即可，[传送门](https://docs.mongodb.com/manual/core/authorization/)

### 开启访问控制

如前面所述，在配置文件中取消`security`注释，在其下面添加`authorization: enabled`

```yaml
security:
  authorization: enabled
```

然后`systemctl restart mongod`重启mongodb即可生效

### 登录

todo





ref

- https://docs.mongodb.com/manual/introduction/