package member

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mongmx/system-d/application/domain/auth"
)

func memberAllEndpoint(s Service, as auth.Service) gin.HandlerFunc {
	type response struct {
		Message string  `json:"message"`
		Member  *Member `json:"member"`
	}
	return func(c *gin.Context) {
		p, err := as.Authorize("user:1", "domain:system", "user.profile.read")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if !p {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden to read profile"})
			c.Abort()
			return
		}
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

func memberGetEndpoint(s Service, as auth.Service) gin.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}
	return func(c *gin.Context) {
		p, err := as.Authorize("1", "system", "user.profile.read")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if !p {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden to read profile"})
			c.Abort()
			return
		}
		id := c.Param("id")
		_, err = s.FindMember(id)
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

func memberListEndpoint(s Service, as auth.Service) gin.HandlerFunc {
	type request struct {
		ID string `json:"id" binding:"required"`
	}
	type response struct {
		Message string `json:"message"`
	}
	return func(c *gin.Context) {
		p, err := as.Authorize("user:1", "domain:system", "user.profile.read")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if !p {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden to read profile"})
			c.Abort()
			return
		}
		var req request
		err = c.ShouldBindJSON(&req)
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
