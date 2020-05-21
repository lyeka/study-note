# Go语言基础



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

切片即动态数组，长度不是固定的，故长度并不是切片的类型组成部分

```go
type SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}
```



**空切片和nil切片**

```go
func main() {
	var s1 []int
	var s2 = []int{}

	fmt.Println("s1:", s1==nil, len(s1), cap(s1))
	fmt.Println("s2:", s2==nil, len(s2), cap(s2))
}
// output
s1: true 0 0
s2: false 0 0
```

上述例子中，s1为nil切片，s2为空切片，虽然他们都长度和容量都是0，但nil切片底层的数据指针为nil。



**遍历**

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



**添加元素**

在容量不足的情况下，append操作会导致重新分配内存



尾部追加元素

```go
var a []int
a = append(a, 1)               // 追加1个元素
a = append(a, 1, 2, 3)         // 追加多个元素, 手写解包方式
a = append(a, []int{1,2,3}...) // 追加一个切片, 切片需要解包
```



开头添加元素

```go
var a = []int{1,2,3}
a = append([]int{0}, a...)        // 在开头添加1个元素
a = append([]int{-3,-2,-1}, a...) // 在开头添加1个切片
```



中间添加元素

```go
var a []int
a = append(a[:i], append([]int{x}, a[i:]...)...)     // 在第i个位置插入x
a = append(a[:i], append([]int{1,2,3}, a[i:]...)...) // 在第i个位置插入切片


// 避免创建中间临时切片
// 插入单个元素
a = append(a, 0)     // 切片扩展1个空间
copy(a[i+1:], a[i:]) // a[i:]向后移动1个位置
a[i] = x             // 设置新添加的元素

// 插入多个元素
a = append(a, x...)       // 为x切片扩展足够的空间
copy(a[i+len(x):], a[i:]) // a[i:]向后移动len(x)个位置
copy(a[i:], x)            // 复制新添加的切片
```



**删除元素**

删除尾部元素

```go
a = []int{1, 2, 3}
a = a[:len(a)-1]   // 删除尾部1个元素
a = a[:len(a)-N]   // 删除尾部N个元素
```



删除头部元素

```go
a = []int{1, 2, 3}
a = a[1:] // 删除开头1个元素
a = a[N:] // 删除开头N个元素

a = []int{1, 2, 3}
a = append(a[:0], a[1:]...) // 删除开头1个元素
a = append(a[:0], a[N:]...) // 删除开头N个元素

a = []int{1, 2, 3}
a = a[:copy(a, a[1:])] // 删除开头1个元素
a = a[:copy(a, a[N:])] // 删除开头N个元素
```

删除中间元素

```go
a = []int{1, 2, 3, ...}

a = append(a[:i], a[i+1:]...) // 删除中间1个元素
a = append(a[:i], a[i+N:]...) // 删除中间N个元素

a = a[:i+copy(a[i:], a[i+1:])]  // 删除中间1个元素
a = a[:i+copy(a[i:], a[i+N:])]  // 删除中间N个元素
```

**切片内存技巧**

巧用空切片

```go
// 删除字节数组中的空格
func TrimSpace(s []byte) []byte {
    b := s[:0]
    for _, x := range s {
        if x != ' ' {
            b = append(b, x)
        }
    }
    return b
}
```

## 函数， 方法和接口

![](https://chai2010.cn/advanced-go-programming-book/images/ch1-11-init.ditaa.png)

上图为Go包初始化流程。一个包内可以有多个`init`函数， `init`函数在同一份文件的话，按顺序执行，不同文件相同包内的`init`函数执行顺序未定义（可能按文件名顺序）。

Go程序的初始化和执行从main包的main函数开始，在其执行之前所有代码都在同一个`goroutine`，也就是程序的主系统线程，因此，某个`init`函数内启动新`goroutine`话，新的`goroutine`要在进入`main.main`函数后才能执行到。



### 函数

函数包括具名函数和匿名函数，支持多个参数和多个返回值。参数还支持可变数量的参数（可以用来当成可选函数？）， 其必选在参数列表最后出现，实际上可变数量函数为切片。

```go
// 多个参数和多个返回值
func Swap(a, b int) (int, int) {
    return b, a
}

// 可变数量的参数
// more 对应 []int 切片类型
func Sum(a int, more ...int) int {
    for _, v := range more {
        a += v
    }
    return a
}
```

当可变参数是一个空接口类型时，调用者是否解包可变参数会导致不同的结果：

```go
func main() {
    var a = []interface{}{123, "abc"}

    Print(a...) // 123 abc
    Print(a)    // [123 abc]
}

func Print(a ...interface{}) {
    fmt.Println(a...)
}
```

需要注意的话，虽然可变数量参数为`...interface{}`

但是直接将具体类型切片解包赋值给`...interface{}`是会保错的

```go
func main() {
	var a = []int{123, 246}

	Print(a...) // 报错 cannot use a (type []int) as type []interface {} in argument to Print
	Print(a)    // 正常
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}
```

因为`...interface{}`本质是`[]interface{}`, 解包赋值等同于将`[]int`转换为`[]interface{}`，这是不被允许的（详见https://github.com/golang/go/wiki/InterfaceSlice），而第二种不解包直接赋值等同于`[]interface{}[0] = a`，是可以正常通过的。



返回值可以被命名，如果返回值命名了，可以通过名字来修改返回值，也可以通过`defer`语句在`return`语句之后修改返回值：

```go
func Inc() (v int) {
    defer func(){ v++ } ()
    return 42 
}
```

`return 42`等同于 `v = 42; return v`

不过这个函数最终返回的是43， 因为使用`defer`延迟了一个匿名函数，这个函数捕获了外部函数v ，形成了闭包函数，闭包对外部变量的访问是引用方式，故可以改变v的值。

注意，延迟函数是会在`return`之前（不包含return后面的语句）执行的。上述例子实际上的顺序为

1. v = 42
2. 执行 defer 函数
3. return

因为闭包是以引用来访问变量的话，故可能会导致一些隐含问题

```go
func main() {
    for i := 0; i < 3; i++ {
        defer func(){ println(i) } ()
    }
}
// Output:
// 3
// 3
// 3
```

上面每个defer 函数以引用了i变量，在return之前执行，这时候循环已经完毕，i最终为3，所以三个输出都为3

解决办法

```go
func main() {
    for i := 0; i < 3; i++ {
        i := i // 定义一个循环体内局部变量i
        defer func(){ println(i) } ()
    }
}

func main() {
    for i := 0; i < 3; i++ {
        // 通过函数传入i
        // defer 语句会马上对调用参数求值
        defer func(i int){ println(i) } (i)
    }
}
```

函数是是以传值的方式来访问传参数的，但如果参数本身为指针或者引用类型（如切片等），还是可以通过指针改变原参数的值的。

思考： 为什么append必须返回一个切片而不是直接改变切片呢？因为函数是传值的，虽然通过数据指针改变了底层的数据，但是len字段个cap字段没法改变，所以还是得通过返回新切片来更新切片。



Go中的函数递归调用在深度逻辑上没有限制（物理限制有），Go会动态的调整栈的大小。

在Go1.4后，Go的动态栈在内存上改为连续了，当所需内存不够时，runtime会自动将数据搬移到新的内存空间，栈变量的指针也会被自动更新，所以要记住Go的指针不是固定不变的，虽然对外使用无影响，但用一个数值变量保存指针的方式来访问指针指向会不靠谱。

Go中模糊了栈和堆的概念，使用者无须知道局部变量是保存在栈还是堆中，只需要知道他们可以正常工作即可，示例略。



### 方法

方法就是关联到类型的函数。



**面对对象**

首先需要了解一下面对对象三大特征

- 封装
- 继承
- 多态

关于多态可以阅读https://zhuanlan.zhihu.com/p/37655397



Go是不是面对对象语言不重要，Go官方给出的答案是yes and no，只能说Go支持面对对象编程，因为它提供了面对对象的三大特征Go均可以实现

封装就不啰嗦了，好理解。

关于继承，Go中是通过组合来实现的，通过组合，可以继承成员的内部成员以及成员的方法。

```go
type Cache struct {
    m map[string]string
    sync.Mutex
}

func (p *Cache) Lookup(key string) string {
    p.Lock()
    defer p.Unlock()

    return p.m[key]
}
```

`Cache`这个结构体内部嵌入了`sync.Mutex`,  也就继承了`sync.Mutex`的`Lock`以及`Unlock`方法，虽然可以直接通过`p.Lock()`/`p.Unlock()`调用，但是实际上`lock`/`Unlock`方法的接收者为`Cache`的`sync.Mutex`成员。这种展开是编译期间完成的，没有运行时代价。

Go语言中方法是编译时静态绑定的。如果需要多态特性，我们需要借助Go语言接口来实现。



### 接口

Go在提供严格类型检查的同时，通过接口类型实现了对鸭子类型的支持。接口类型是对其它类型行为的抽象和概括，Go中的接口类型独特支出在于它是满足隐式实现的鸭子类型。

Go语言中，对于基础类型（非接口类型）不支持隐式的转换，我们无法将一个`int`类型的值直接赋值给`int64`类型的变量，也无法将`int`类型的值赋值给底层是`int`类型的新定义命名类型的变量。Go语言对基础类型的类型一致性要求可谓是非常的严格，但是Go语言对于接口类型的转换则非常的灵活。对象和接口之间的转换、接口和接口之间的转换都可能是隐式的转换。可以看下面的例子：

```go
var (
    a io.ReadCloser = (*os.File)(f) // 隐式转换, *os.File 满足 io.ReadCloser 接口
    b io.Reader     = a             // 隐式转换, io.ReadCloser 满足 io.Reader 接口
    c io.Closer     = a             // 隐式转换, io.ReadCloser 满足 io.Closer 接口
    d io.Reader     = c.(io.Reader) // 显式转换, io.Closer 不满足 io.Reader 接口
)
```

[todo]