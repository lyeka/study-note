# ctrl-c ctrl-z ctrl-d 区别

在Linux中：



ctrl-c: ( kill foreground process ) 强制终止当前进程

ctrl-z: ( suspend foreground process ) 挂起当前进程，可以使用`fg/bg`命令恢复进程，shell关闭后所有挂起的任务会被强制终止

ctrl-d: ( Terminate input, or exit shell ) 一个特殊的二进制值，表示 EOF，作用相当于在终端中输入exit后回车；





ref

- [Linux中ctrl-c, ctrl-z, ctrl-d 区别](https://blog.csdn.net/mylizh/article/details/38385739)