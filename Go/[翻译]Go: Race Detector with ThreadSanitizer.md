# Go: Race Detector with ThreadSanitizer

Go 提供了强大的工具用于检测竞态。在测试或者编译时可以使用 `-race` 标志来启动该工具。

让我们来创建一个简单的数据竞争的例子来分析 Go 如何以及使用什么工具来作竞态检测。下面程序将通过两个 goroutines 来增加一个变量到十万。

```go
var foo = 0

func main() {
    var wg sync.WaitGroup

    wg.Add(2)
    go func() {
        defer wg.Done()
        for i := 0; i < 50000 ;i++  {
            foo++
        }
    }()
    go func() {
        defer wg.Done()
        for i := 0; i < 50000 ;i++  {
            foo++
        }
    }()
    wg.Wait()
    println(foo)
}
```

程序每次输出的结果都会因为并发读写该变量而不一样。如果你不能完全理解该例子，建议先去阅读 [Introducing the Go Race Detector](https://blog.golang.org/race-detector) 。

## 竞态检测功能

为了更改地理解 Go 内部如何管理竞态检测， 我们先使用 `go tool compile -S -race main.go` 命令来生成  [asm](https://golang.org/doc/asm) 。下面是与 `foo++` 操作相关的输出摘录：

```bash
(main.go:14)   CALL   runtime.raceread(SB)
(main.go:14)   CALL   runtime.racewrite(SB)
(main.go:16)   CALL   runtime.racefuncexit(SB)
[...]
(main.go:20)   CALL   runtime.raceread(SB)
(main.go:20)   CALL   runtime.racewrite(SB)
(main.go:22)   CALL   runtime.racefuncexit(SB)
```

因为 `foo++` 操作等同于 `foo = foo + 1` ，Go 必须在改变其值前先读取该变量。Go 在读写该变量期间添加了两个函数用于观察该变量上是否存在竞态。我们到 `runtime` 包查看下这些函数是什么。

## 竞态检测包

Go 在 `runtime` 包提供了两份与竞态检测相关的文件：`race.go` 和 `race0.go` ：

- `race.go` 设置 `raceenable` 常量为 true 以及提供与竞态检测相关的方法：

```go
package runtime

const raceenabled = true

func raceread(uintptr)
func racewrite(uintptr)
[...]
```



- `race0.go` 设置 `receenable` 常量为 false 并且提供了相同的方法，但是函数体不同，会抛出错误。

```go
package runtime

const raceenabled = false

func raceacquire(addr unsafe.Pointer) { throw("race") }
func raceacquireg(gp *g, addr unsafe.Pointer) { throw("race") }
[...]
```

当程序不应该使用竞态检测时，将会抛出错误来中止竞态检测。

`raceenable` 常量将会在 Go 库 主要是 `runtime` 包中使用，用在程序运行时添加特定用于竞态检测的检查点。



是否导入 `runtime/race` 包在 `load` 包中的 `internal/load/pkg.go` 中决定

```go
// LinkerDeps returns the list of linker-induced dependencies for main package p.
func LinkerDeps(p *Package) []string {
    [...]
    // Using the race detector forces an import of runtime/race.
    if cfg.BuildRace {
        deps = append(deps, "runtime/race")
    }

    return deps
}
```

如我们所见， `-race` 标志映射到内部配置 `cfg` 中。



## 数据竞态检测标志

Go 在`cfg` 包中保留了内部配置用于映射所有的标志。 下面是配置示例：

```go
package cfg

// These are general "build flags" used by build and other commands.
var (
    [...]
    BuildP                 = runtime.NumCPU() // -p flag
    BuildRace              bool               // -race flag=
    BuildV                 bool // -v flag
    [...]
```

使用数据竞态标志 `-race` 将会[更新此内部配置](https://github.com/golang/go/blob/release-branch.go1.12/src/cmd/go/internal/work/build.go#L228)：

```go
build.go
// addBuildFlags adds the flags common to the build, clean, get,
// install, list, run, and test commands.
func AddBuildFlags(cmd *base.Command) {
   [...]
   cmd.Flag.BoolVar(&cfg.BuildRace, "race", false, "")
   [...]
}
```

如果想排除不想用于竞态检测的文件，可以使用 `// +build !race` 标签(注释)。

在对数据竞态检测器内部的工作流程有了更好的了解后，让我们回到最初，理解这些方法如何工作。

## ThreadSanitizer

在内部，`raceread` 方法会将控制委派给另外一个方法：

```go
// func runtime·raceread(addr uintptr)
// Called from instrumented code.
TEXT   runtime·raceread(SB), NOSPLIT, $0-8
   MOVQ   addr+0(FP), RARG1
   MOVQ   (SP), RARG2
   // void __tsan_read(ThreadState *thr, void *addr, void *pc);
   MOVQ   $__tsan_read(SB), AX
   JMP    racecalladdr<>(SB)
```

该方法属于  [ThreadSanitizer](http://clang.llvm.org/docs/ThreadSanitizer.html) （又名 Tsan）数据竞态检测工具的一部分，如文档所示，开启该工具会使程序运行变慢，故不建议在生产环境中使用。

> 使用 *ThreadSanitizer* 会使运行速度下降约5到15倍，内存开销增加约5到10倍。

在第一个例子中， `raceread` 和 `racewrite` 方法取得了 `foo` 变量的内存地址读写权限以检测是否有潜在的数据竞争发生。

如果想更深入地了解 ThreadSanitizer ， 可以去阅读 [Kavya Joshi](http://kavya joshi channels/)  的 “[Looking Inside a Race Detector](https://www.infoq.com/presentations/go-race-detector/)” 。



```
---
via: https://medium.com/a-journey-with-go/go-race-detector-with-threadsanitizer-8e497f9e42db

作者：[Vincent Blanchon](https://medium.com/@blanchon.vincent)
译者：[译者ID](https://github.com/译者ID)
校对：[校对者ID](https://github.com/校对者ID)

本文由 [GCTT](https://github.com/studygolang/GCTT) 原创编译，[Go 中文网](https://studygolang.com/) 荣誉推出
```