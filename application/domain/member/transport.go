package member

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var s Service

// Routes in the domain route
func Routes(route *gin.Engine, service Service) {
	s = service
	r := route.Group("/member/v1/admin")
	{
		r.GET("/lists", memberAllEndpoint(s))
		r.GET("/list/:id", memberGetEndpoint(s))
		r.POST("/list", mustLogin(), memberListEndpoint(s))
	}
}

func mustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if true {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user needs to be signed in to access this service"})
			c.Abort()
			return
		}
		c.Next()
	}
}
