package mock

import "github.com/gin-gonic/gin"

func Authorizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
