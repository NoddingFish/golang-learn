## Go 使用 Redis

#### 前言

> [go操作redis](https://www.cnblogs.com/gwyy/p/13289126.html)

`golang` 操作 `redis` 主要有两个库， `go-redis` 和 `redigo` 。两者操作都比较简单，区别上 `redigo` 更像一个 `client` 执行各种操作都是通过 `Do` 函数去做的， `redisgo` 对函数的封装更好，相比之下 `redigo` 操作 `redis` 显得有些繁琐。但是官方更推荐 `redigo` ，所以项目中我使用了 `redigo` 。

`redigo` 的使用入门可以去查 `godoc` ：http://godoc.org/github.com/garyburd/redigo/redis



#### 安装第三方 `redis` 库

​	在 `GOPATH` 路径下执行安装指令：

```shell
go get github.com/garyburd/redigo/redis
```