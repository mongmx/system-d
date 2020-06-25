package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var s Service

// Routes in the domain route
func Routes(route *gin.Engine, service Service) {
	s = service
	r := route.Group("/auth/v1")
	{
		r.GET("/login", loginEndpoint(s))
		r.GET("/token/refresh", refreshTokenEndpoint(s))
		r.POST("/token/revoke/:id", mustLogin(), revokeTokenEndpoint(s))
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
