# runtime 运行时

## 什么是 runtime

在计算机领域中，经常会接触到 runtime 这个概念，那么 runtime  究竟是什么东西？

runtime 描述了**程序运行时候**执行的软件/指令， 在每种语言有着不同的实现。可大可小，在 C 中，runtime 是库代码， 等同于 C runtime library，一系列 C 程序运行所需的函数，在 Java 中，runtime 还提供了 Java 程序运行所需的虚拟机等。

总而言之，**runtime 是一个通用抽象的术语，指的是计算机程序运行的时候所需要的一切代码库，框架，平台等**。



## Go中的 runtime

在 Go 中，  有一个 runtime 库，其实现了垃圾回收，并发控制， 栈管理以及其他一些 Go 语言的关键特性。 runtime 库是每个 Go 程序的一部分，也就是说编译 Go 代码为机器代码时也会将其也编译进来。所以 Go 官方将其定位偏向类似于 C 语言中的库。Go 中的 runtime 不像 Java runtime （JRE， java runtime envirement ) 一样，jre 还会提供虚拟机， Java 程序要在 JRE 下 才能运行。

所以在 Go 语言中， runtime 只是提供支持语言特性的库的名称，也就是 Go 程序执行时候使用的库。



ref 

- [what-is-runtime](https://stackoverflow.com/questions/3900549/what-is-runtime)

- https://golang.org/doc/faq 中 【 Does Go have a runtime? 】 部分

- [GopherCon 2019 - Controlling the go runtime](https://about.sourcegraph.com/go/gophercon-2019-controlling-the-go-runtime)

    