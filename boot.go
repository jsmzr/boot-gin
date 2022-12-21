package boot

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	plugin "github.com/jsmzr/boot-plugin"
	"github.com/spf13/viper"
)

const configPrefix = "boot.gin"

var routerInitFuncions []func(*gin.Engine)

var defaultConfig map[string]interface{} = map[string]interface{}{"port": 8080}

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
	return engine.Run(fmt.Sprintf(":%d", viper.GetInt(configPrefix+".port")))
}

func init() {
	for key := range defaultConfig {
		viper.SetDefault(configPrefix+"."+key, defaultConfig[key])
	}
}
