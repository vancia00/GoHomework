package middleware

import (
	"GoHomework/tools/JwtToken"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "请求头中auth为空",
			})
			c.Abort()
			return
		}
		//按空格分隔
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		//parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := JwtToken.ValidateRegisteredClaim(parts[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "无效的Token",
			})
			c.Abort()
			return
		}
		//将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Next()
	}
}
