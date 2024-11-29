package handler

import (
	"authentication-service/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var creds domain.Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	tokenString, err := h.services.Auth.Login(creds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (h *Handler) Authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.Abort()
		return
	}

	username, err := h.services.Auth.Auth(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("username", username)
	c.Next()
}

func (h *Handler) ProtectedEndpoint(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{"message": "Valid token, welcome " + username})
}
