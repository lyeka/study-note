# Go基础

## 数组，字符串与切片

Go语言中，数组，字符串，切片三种类型紧密相关，它们在底层原始数据有着相同的内存结构，在上层限制了语法而有着不同的表行。数组类型是字符串与切片类型的基础，关于数组的很多操作可以直接用于字符串与切片中。

### 数组

数组是由**固定长度**的**相同类型元素**组成的序列。长度，元素类型都是数组的组成部分，不同长度，相同类型的数组是不同的类型，并不可以直接赋值、比较。



声明与初始化数组

```go
var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
var b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6
```

在Go中，数组是值语义，一个数组变量即表示整个数组（不像C中数组变量是隐式指向数组第一个元素的指针），在数组变量被赋值或者被传递的时候，会复制整个数组，可以通过传递指向数组指针来规避这个开销。



通过数组指针和数组来操作数组在语法上并没有差别

```go
var a = [...]int{1, 2, 3} // a 是一个数组
var b = &a                // b 是指向数组的指针

fmt.Println(a[0], a[1])   // 打印数组的前2个元素
fmt.Println(b[0], b[1])   // 通过数组指针访问数组元素的方式和数组类似

for i, v := range b {     // 通过数组指针迭代数组的元素
    fmt.Println(i, v)
}
```



在数组中，len(长度)与cap(容量)是一致的。

迭代数组的几种方式，其中使用`for range`方式可以规避数组越界问题。

```go
for i := range a {
    fmt.Printf("a[%d]: %d\n", i, a[i])
} 
for i, v := range b {
    fmt.Printf("b[%d]: %d\n", i, v)
}
for i := 0; i < len(c); i++ {
    fmt.Printf("c[%d]: %d\n", i, c[i])
}
```



### 字符串

字符串底层数据其实就是字节数组，不过只只读的，也就是是不可以改变底层的字节元素。字符串的长度虽然是固定的，但是字符串的长度并不是字符串类型的一部分。

字符串的底层结构在`reflect.StringHeader`中定义

```go
type StringHeader struct {
    Data uintptr
    Len  int
}
```

Data字段即指向底层字节数组的指针，Len字段表示字符串的长度。



字符串支持切片操作，取出来的是字符串类型，通过位置索引取的话是byte/uint8类型

```go
s := "hello"
fmt.Println(reflect.TypeOf(s))  // string
fmt.Println(reflect.TypeOf(s[0])) // uint8
fmt.Println(reflect.TypeOf(s[1:])) // string
```



Go语言源文件采用utf8编码，其中出现的字符串字面量一般也是utf8编码（对于转义字符择没有这个限制）。

```go
fmt.Printf("%#v\n", []byte("Hello, 世界"))

// output
[]byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c}
```

因为utf8是变长编码，所有底层字节数组的长度个字面量字符数量并不一定一致，上述例子中，中文在utf8中的是用3个字节表示的，所以字节数组的长度大于字面量字符数量。



迭代字符串不同的方法也会出现不同的结果

一、数组迭代

```go
s := "w噢"
fmt.Println(reflect.TypeOf(s[1]))
for i:=0; i<len(s); i++{
    fmt.Println(s[i])
    fmt.Println(string(s[i]))
}

// output
// uint8
// src:  119 after string:  w
// src:  229 after string:  å
// src:  153 after string:  
// src:  162 after string:  ¢
```

这种方式是直接迭代了字符串底层的字节数组



二、`for range` 迭代

```go
i := 1
for _, v := range s {
    if i == 1 {
        fmt.Println(reflect.TypeOf(v))
        i+=1
    }
    fmt.Println("src:", v, "after string: ", string(v))
}

// output
// int32
// src: 119 after string:  w
// src: 22114 after string:  噢
```

使用`for range`迭代的话会遍历utf8解码后的unicode码点值



如何得到字符串中的字符数量

```go
fmt.Println(len("w噢"))
fmt.Println(utf8.RuneCountInString("w噢"))
fmt.Println(len([]rune("w噢")))

// output
// 4
// 2
// 2
```



### 切片

todo

