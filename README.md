# boot-gin

[![Build Status](https://github.com/jsmzr/boot-gin/workflows/Run%20Tests/badge.svg?branch=main)](https://github.com/jsmzr/boot-gin/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/jsmzr/boot-gin/branch/main/graph/badge.svg?token=HNQCAN3UVR)](https://codecov.io/gh/jsmzr/boot-gin)

提供 gin 框架的 boot 支持

## 使用说明

项目运行，[详细示例](https://github.com/jsmzr/gin-boot-example)

```go
package main

import (
	"fmt"

	boot "github.com/jsmzr/boot-gin"
    // 通过声明的方式引入需要使用的插件
	_ "github.com/jsmzr/boot-plugin-config-yaml"
	_ "github.com/jsmzr/boot-plugin-logrus"
	_ "github.com/jsmzr/boot-plugin-prometheus"
	_ "github.com/jsmzr/gin-boot-example/demo"
	_ "github.com/jsmzr/gin-boot-example/middlewares"
	_ "github.com/jsmzr/gin-boot-example/user"
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
