package Public

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuthCheck() func(c *gin.Context){
	return func(c *gin.Context) {
		userid := c.MustGet("Userid").(int64)
		username := c.MustGet("Username").(string)
		if username != "" && userid != 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"msg":  "success",
				"data": gin.H{
					"userid":   userid,
					"username": username,
				},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 2001,
				"msg":  "error",
				"data": "JWT为空或者错误",
			})
		}
	}
}

func EncryptPassword(Password string) string {
	h := md5.New()
	h.Write([]byte("logan.com"))
	return hex.EncodeToString(h.Sum([]byte(Password)))
}
