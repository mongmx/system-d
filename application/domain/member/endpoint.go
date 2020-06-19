package member

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func memberAllEndpoint(s Service) gin.HandlerFunc {
	type response struct {
		Message string  `json:"message"`
		Member  *Member `json:"member"`
	}
	return func(c *gin.Context) {
		member, err := s.FindAllMember()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, &response{
			Message: "success",
			Member:  member,
		})
	}
}

func memberGetEndpoint(s Service) gin.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}
	return func(c *gin.Context) {
		id := c.Param("id")
		_, err := s.FindMember(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, &response{
			Message: "success",
		})
	}
}

func memberListEndpoint(s Service) gin.HandlerFunc {
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
		_, err = s.FindMember(req.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, &response{
			Message: "success",
		})
	}
}
