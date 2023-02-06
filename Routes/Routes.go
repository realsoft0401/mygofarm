package Routes

import (
	"fmt"
	"mygofarm/Middlewares"
	"mygofarm/Pkg/Logger"
	"mygofarm/Services/User"
	"mygofarm/Settings"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.Use(Logger.GinLogger(), Logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("系统版本信息:%s\n", Settings.Conf.Version))
	})
	r.POST("/signadd", User.SignHandler)
	r.POST("/login", User.LoginHandler)
	r.POST("/signdel", Middlewares.JWTAuthMiddleware(),User.SignDelHandler)

	return r
}
