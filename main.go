package main

import (
	"scaffold-demo/config"
	// _ "scaffold-demo/config" 会自动调用init函数而不需要手动使用里面的方法
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	logs.Info(nil, "启动服务器...")
	// logs.Debug(nil, "这里是debug...")
	// logs.Warn(nil, "这里是warn...")
	// logs.Error(nil, "这里是error...")
	port := config.Port
	r.Run(port)
}
