# Shell、 sh、Bash的联系

## Shell

在Linux中，**Shell是一个应用程序**，是用户与系统内核kernel之间的代理，命令行键入的命令会通过Shell传递给内核以操纵计算机资源。

另外**Shell也是一种编程语言**，其编译器就是Shell，Shell脚本就是使用Shell语言写的程序脚本



最初Linux是只有Shell这个工具的，系统启动后就进入控制台（Console）就是所谓的黑框，但是随着图形桌面化的发展，Linux（桌面版本）启动后是直接进入了桌面程序，而不是控制台。现代Linux启动后会创建一些虚拟控制台（Virtual Console)， 可以通过`Ctrl+Alt+Fn(n=1,2,3,4,5...)`进入到不同的虚拟控制台，桌面化程序就是运行在第一个虚拟控制台`Ctrl+Alt+F1`。如果想要进入黑框的虚拟控制台，可以按`Ctrl+Alt+Fn(n=2,3,4,5...)`，想回到桌面按`Ctrl+Alt+F1`即可



但是在现在Linux(桌面版本)中，一般通过终端（Terminal）来使用Shell。

## sh、Bash

Bash其实就是Shell的一种。

Shell作为一个应用程序不仅仅只有一种，常见的Shell有sh， Bash， csh等，在这个维度上Shell的概念是指用户与内核交互的一类应用程序。

sh(/bin/sh)是最早的Shell，Bash(/bin/bash)是sh的的扩展，兼容sh，但是比sh多了许多功能（如命令补全等），Bash是现在最流行的shell，大部分Linux的默认的Shell就是Bash。

在一般情况下，人们并不区分 sh 和 Bash，所以，像 `#!/bin/sh`，它同样也可以改为 `#!/bin/bash`。

`#!` 告诉系统其后路径所指定的程序即是解释此脚本文件的 Shell 程序。