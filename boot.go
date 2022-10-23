package boot

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	config "github.com/jsmzr/boot-config"
	plugin "github.com/jsmzr/boot-plugin"
)

type GinProperties struct {
	Port *int
	// TODO many
}

var defaultPort = 8080
var routerInitFuncions []func(*gin.Engine)

func log(message string) {
	fmt.Printf("[BOOT-GIN] %v| %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}

func RegisterRouter(f func(*gin.Engine)) {
	routerInitFuncions = append(routerInitFuncions, f)
}

func Run() error {
	if err := plugin.PostProccess(); err != nil {
		return err
	}
	var properteis GinProperties
	_ = config.Resolve("boot.gin", &properteis)
	if properteis.Port == nil {
		properteis.Port = &defaultPort
	}
	engine := gin.Default()
	log("init gin middleware")
	if err := initMiddleware(engine); err != nil {
		return err
	}
	log("init gin router")
	for _, f := range routerInitFuncions {
		f(engine)
	}
	log("init gin service")
	return engine.Run(fmt.Sprintf(":%d", *properteis.Port))
}
