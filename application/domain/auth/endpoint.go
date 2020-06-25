package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginEndpoint(s Service) gin.HandlerFunc {
	type response struct {
		Message string  `json:"message"`
	}
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, &response{
			Message: "success",
		})
	}
}

func refreshTokenEndpoint(s Service) gin.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, &response{
			Message: "success",
		})
	}
}

func revokeTokenEndpoint(s Service) gin.HandlerFunc {
	type request struct {
		ID string `json:"id" binding:"required"`
	}
	type response struct {
		Message string `json:"message"`
	}
	return func(c *gin.Context) {
		var req request
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request", "request": req})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, &response{
			Message: "success",
		})
	}
}
