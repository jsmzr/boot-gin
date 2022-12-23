package boot

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

type TestGinMiddleware struct{}
type TestErrorGinMiddleware struct{}

func (m *TestGinMiddleware) Load(e *gin.Engine) error {
	return nil
}

func (m *TestGinMiddleware) Order() int {
	return 0
}
func (m *TestErrorGinMiddleware) Load(e *gin.Engine) error {
	return fmt.Errorf("load middle error")
}

func (m *TestErrorGinMiddleware) Order() int {
	return 1
}

func TestRegisterMiddleware(t *testing.T) {
	middlewares = make(map[string]GinMiddleware)
	name := "test"
	RegisterMiddleware(name, &TestGinMiddleware{})
	if _, ok := middlewares[name]; !ok {
		t.Fatal("register failed")
	}
	defer func() {
		if err := recover(); err == nil {
			t.Fatal("middle already exists")
		}
	}()
	RegisterMiddleware(name, &TestGinMiddleware{})
}

func TestInitMiddleware(t *testing.T) {
	middlewares = make(map[string]GinMiddleware)
	// len == 0
	engin := gin.Default()
	if err := initMiddleware(engin); err != nil {
		t.Fatal(err)
	}

	RegisterMiddleware("test", &TestGinMiddleware{})
	// success
	if err := initMiddleware(engin); err != nil {
		t.Fatal(err)
	}

	RegisterMiddleware("testError", &TestErrorGinMiddleware{})
	if err := initMiddleware(engin); err == nil {
		t.Fatal("init should be error")
	}

}
