# boot-gin

[![Build Status](https://github.com/jsmzr/boot-gin/workflows/Run%20Tests/badge.svg?branch=main)](https://github.com/jsmzr/boot-gin/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/jsmzr/boot-gin/branch/main/graph/badge.svg?token=HNQCAN3UVR)](https://codecov.io/gh/jsmzr/boot-gin)

gin 的 boot 适配库，灵活的使用 boot 系列库的能力，并简化 gin 路由和全局中间件使用。遵循约定大于配置，简化项目的搭建与使用。

- 约定大于配置，插件+配置的方式直接使用，降低接入成本
- 路由注册，简化路由管理
- 全局中间件，简化全局中间件及第三方中间件使用

## 使用说明

项目运行，[详细示例](https://github.com/jsmzr/gin-boot-example)。

```go
package main

import (
	"fmt"

	boot "github.com/jsmzr/boot-gin"
    // boot 系列插件的注册
	_ "github.com/jsmzr/boot-plugin-logrus"
	// 路由注册
	_ "github.com/jsmzr/gin-boot-example/router"
	// 全局中间件注册
	_ "github.com/jsmzr/gin-boot-example/middlewares"
)

func main() {
	if err := boot.Run(); err != nil {
		fmt.Println(err)
	}
}

```

路由注册

```go
// 路由注册，通过 init 方法注册路由
// router.go
func InitRouter(e *gin.Engine) {
    e.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "pong"})
	})
    e.GET("/foo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "bar"})
	})
}
func init() {
    boot.RegisterRouter(InitRouter)
}
```

中间件注册

```go
// 中间件注册与声明，需要实现 GinMiddleware 编写 load 逻辑及中间件加载顺序，通过 init 注册
type DemoMiddleware struct{}

func (d *DemoMiddleware) Load(e *gin.Engine) error {
	e.Use(func(ctx *gin.Context) {
		log.Info("DemoMiddleware start")
		ctx.Next()
		log.Info("DemoMiddleware end")
	})
	return nil
}

func (d *DemoMiddleware) Order() int {
	return 0
}

func init() {
	boot.RegisterMiddleware("demo", &DemoMiddleware{})
}
```
