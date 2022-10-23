package boot

import (
	"testing"

	"github.com/gin-gonic/gin"
)

type TestGinMiddleware struct{}

func (m *TestGinMiddleware) Load(e *gin.Engine) error {
	return nil
}

func (m *TestGinMiddleware) Order() int {
	return 0
}

func TestRegisterMiddleware(t *testing.T) {
	middlewares = make(map[string]GinMiddleware)
	name := "test"
	RegisterMiddleware(name, &TestGinMiddleware{})
	if _, ok := middlewares[name]; !ok {
		t.Fatal("register failed")
	}
}
