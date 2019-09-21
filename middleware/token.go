package middleware

import "github.com/gin-gonic/gin"

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("doumi", 999)
		c.Next()
		return
	}
}
