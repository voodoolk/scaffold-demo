package auth

import (
	authController "scaffold-demo/controllers/auth"

	"github.com/gin-gonic/gin"
)

func routeLogin(authGroup *gin.RouterGroup) {
	authGroup.POST("/login", authController.ApiLogin)
}

func routeLogout(authGroup *gin.RouterGroup) {
	authGroup.GET("/logout", authController.ApiLogout)
}

func RegisterSubRoutes(g *gin.RouterGroup) {
	authGroup := g.Group("/auth")
	//传给路由处理函数
	routeLogin(authGroup)
	routeLogout(authGroup)
}
