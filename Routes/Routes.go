package Routes

import (
	"fmt"
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
	r.POST("/signup", User.SignHandler)

	//r.GET("/ping", Middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
	//
	//	userid := c.MustGet("userid").(int64)
	//	username := c.MustGet("username").(string)
	//	if username != "" && userid != 0 {
	//		c.JSON(http.StatusOK, gin.H{
	//			"code": 2000,
	//			"msg":  "success",
	//			"data": gin.H{
	//				"userid":   userid,
	//				"username": username,
	//			},
	//		})
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{
	//			"code": 2001,
	//			"msg":  "error",
	//			"data": "JWT为空或者错误",
	//		})
	//	}
	//})

	return r
}
