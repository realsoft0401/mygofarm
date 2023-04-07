package Routes

import (
	"fmt"
	"mygofarm/Middlewares"
	"mygofarm/Pkg/Logger"
	"mygofarm/Services/User"
	"mygofarm/Settings"

	"net/http"

	_ "mygofarm/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.Use(Logger.GinLogger(), Logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("欢迎使用 GoFarm 脚手架 \n 系统版本信息:%s\n", Settings.Conf.Version))
	})
	r.POST("/api/signadd", User.SignHandler)
	r.POST("/api/login", User.LoginHandler)
	r.POST("/api/signdel", Middlewares.JWTAuthMiddleware(), User.SignDelHandler)
	r.POST("/api/signone", Middlewares.JWTAuthMiddleware(), User.GetoneUserHandler)

	return r
}
