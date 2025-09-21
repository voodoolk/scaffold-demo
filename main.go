package main

import (
	"fmt"
	"scaffold-demo/config"

	// _ "scaffold-demo/config" 会自动调用init函数而不需要手动使用里面的方法
	"scaffold-demo/utils/jwtutil"
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
	res, _ := jwtutil.GenToken("llll")
	fmt.Println("生成的jwt的token为:", res)
	_, err := jwtutil.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxsbGwiLCJpc3MiOiJsbGtrIiwic3ViIjoic3ViUmFiYml0IiwiZXhwIjoxNzU4NDM4ODAxLCJuYmYiOjE3NTg0MzE2MDEsImlhdCI6MTc1ODQzMTYwMX0.6JUIZdCAjPuqdnyNjuF_drZL1hhqLwtFPr-wEd_ydqk")
	if err != nil {
		fmt.Printf("解析token出错:%v\n", err.Error())
	}
	r.Run(port)
}
