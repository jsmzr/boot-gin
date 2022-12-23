package boot

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	plugin "github.com/jsmzr/boot-plugin"
)

type TestErrorPlugin struct{}

func (t TestErrorPlugin) Enabled() bool {
	return true
}
func (t TestErrorPlugin) Order() int {
	return 1
}
func (t TestErrorPlugin) Load() error {
	return fmt.Errorf("test error")
}

func TestRegisterRouter(t *testing.T) {
	RegisterRouter(func(e *gin.Engine) {})
}
func TestRun(t *testing.T) {
	RegisterMiddleware("testError", &TestErrorGinMiddleware{})
	if Run() == nil {
		t.Fatal("init middleware should be error")
	}
}

func TestRun1(t *testing.T) {
	plugin.Register("testError", &TestErrorPlugin{})
	if Run() == nil {
		t.Fatal("init plugin should be error")
	}
}
