package middleware

import "github.com/gin-gonic/gin"

type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

func (m *GoMiddleware) CORS( /*next *gin.HandlerFunc*/ ) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
