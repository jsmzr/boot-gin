package boot

import (
	"fmt"
	"sort"

	"github.com/gin-gonic/gin"
)

type GinMiddleware interface {
	Load(*gin.Engine) error
	Order() int
}

var middlewares = make(map[string]GinMiddleware)

func RegisterMiddleware(name string, m GinMiddleware) {
	_, ok := middlewares[name]
	if ok {
		panic(fmt.Errorf("gin middleware [%s] already registerd", name))
	}
	log(fmt.Sprintf("Register [%s:%T] middleware", name, m))
	middlewares[name] = m
}

func initMiddleware(e *gin.Engine) error {
	if len(middlewares) == 0 {
		log("Not found Gin Middleware")
		return nil
	}
	values := make([]GinMiddleware, 0, len(middlewares))
	for _, v := range middlewares {
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i].Order() < values[j].Order()
	})
	for i := 0; i < len(values); i++ {
		if err := values[i].Load(e); err != nil {
			return err
		}
		log(fmt.Sprintf("Load [%T] middleware", values[i]))
	}

	return nil
}
