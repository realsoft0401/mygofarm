package Middlewares

import (
	"mygofarm/Pkg/HttpResponse"
	"mygofarm/Pkg/Jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {

			HttpResponse.ResponseError(c, HttpResponse.CodeNotLogin)

			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			HttpResponse.ResponseError(c, HttpResponse.CodeInvalidAuthFormat)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := Jwt.ParseToken(parts[1])
		if err != nil {
			HttpResponse.ResponseError(c, HttpResponse.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("Userid", mc.UserId)
		c.Set("Username", mc.Username)
		c.Next() // 后续的处理函数可以用过c.Get(ContextUserNameIDKey)来获取当前请求的用户信息
	}
}
