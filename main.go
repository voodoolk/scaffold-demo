package main

import (
	_ "scaffold-demo/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Run(":8080")
}
