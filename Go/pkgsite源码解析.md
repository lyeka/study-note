# pkgsite源码解析

[github](https://github.com/golang/pkgsite)

[web](https://pkg.go.dev/)



## 架构



## 代码解析

### 代码目录组织

```
├── all.bash
├── cloudbuild.yaml
├── cmd
│   ├── frontend
│   ├── prober
│   ├── teeproxy
│   └── worker
├── content
│   └── static
├── CONTRIBUTING.md
├── devtools
│   ├── compile_js.sh
│   ├── create_local_db.sh
│   ├── create_migration.sh
│   ├── drop_test_dbs.sh
│   ├── lib.sh
│   └── migrate_db.sh
├── doc
│   ├── architecture.png
│   ├── design.md
│   ├── frontend.md
│   ├── postgres.md
│   ├── precommit.md
│   └── worker.md
├── go.mod
├── go.sum
├── internal
│   ├── auth
│   ├── complete
│   ├── config
│   ├── database
│   ├── datasource.go
│   ├── dcensus
│   ├── derrors
│   ├── discovery.go
│   ├── discovery_test.go
│   ├── experiment
│   ├── experiment.go
│   ├── fetch
│   ├── frontend
│   ├── index
│   ├── licenses
│   ├── log
│   ├── middleware
│   ├── postgres
│   ├── proxy
│   ├── proxydatasource
│   ├── queue
│   ├── secrets
│   ├── source
│   ├── stdlib
│   ├── teeproxy
│   ├── testing
│   ├── version
│   ├── worker
│   └── xcontext
├── LICENSE
├── migrations
│   ├── 000001_initial_schema_from_pg_dump.down.sql
│   ├── ...
│   └── 000021_add_version_map_go_mod_path_column.up.sql
├── PATENTS
├── README.md
└── third_party
    ├── autoComplete.js
    └── dialog-polyfill
```



#### cmd

Go圈内推荐放置main.main的地方。

每个子文件夹为单独一个子项目（可执行文件）

#### internal

该项目内代码，对外项目不可用（因为通常是业务代码）

与此对应的是`pkg`文件夹（通常包含一个通用的工具类代码），`pkg`是对外可的用的，即别的包可以直接import`pkg`

#### devtools

放置一些脚本，方便开发。

如`pkgsite`项目下包含了在本地建测试数据库的脚本

#### doc

存放文档



### 命令行参数

官方的`flag`库即可满足

### 配置

[代码地址](https://github.com/golang/pkgsite/tree/master/internal/config)

- 单例模式，实例化一个Config结构体供读取，配置从环境变量读取。

- Config结构体相关配置可以写在同一行，阅读性好。

- 为什么需要一个ctx？为了从第三方服务获取配置，如密钥服务，所以在初始化配置函数时需要传入一个ctx。

- db密码从第三方服务获取

- 一些组合配置读取可以通过config方法提供，如db dsn
- config不是通过包方法对外提供配置，而是通过实例的方法提供？config实例一直被传递？

### 路由

### 中间件

- 链式调用



### 错误处理、错误码

### 数据库

[代码地址](https://github.com/golang/pkgsite/tree/master/internal/postgres)

- 原生sql
-  单例模式，通过db方法提供对数据的操作
- 

### 日志

[代码地址](https://github.com/golang/pkgsite/tree/master/internal/log)

- 封装不同log级别的方法供外部使用，并提供格式化/非格式化的log
- 为什么(获取)logger需要加锁？
- fatal()级别退出程序
- 两种logger，通过接口的方式提供打日志调用

### 测试

- 







