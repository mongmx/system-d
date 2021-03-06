package member

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mongmx/system-d/application/domain/auth"
)

var s Service

// Routes in the domain route
func Routes(route *gin.Engine, s Service, as auth.Service) {
	r := route.Group("/member/v1/admin")
	{
		r.GET("/lists", memberAllEndpoint(s, as))
		r.GET("/list/:id", memberGetEndpoint(s, as))
		r.POST("/list", mustLogin(), memberListEndpoint(s, as))
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
