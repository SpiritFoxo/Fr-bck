package handlers

import (
	"authservice/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetUserInfo(c *gin.Context) {

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot get user_id"})
		c.Abort()
		return
	}

	user := models.User{}
	result := s.db.First(&user, userId)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot find user"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
	})
}

func (s *Server) UpdateUserInfo(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot get user_id"})
		return
	}

	var user models.User
	if err := s.db.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot find user"})
		return
	}

	var requestBody struct {
		Username string `json:"username"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil || requestBody.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty username"})
		return
	}

	if err := s.db.Model(&user).Update("username", requestBody.Username).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update username"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "username updated successfully"})
}
