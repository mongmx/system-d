package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginEndpoint(s Service) gin.HandlerFunc {
	type request struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	type response struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}
	return func(c *gin.Context) {
		var req request
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request", "request": req})
			c.Abort()
			return
		}
		user, err := s.CheckUser(req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		token, err := s.CreateToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, &response{
			Message: "success",
			Token:   token,
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
