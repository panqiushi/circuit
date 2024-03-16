package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is authenticated
		if _, err := c.Cookie("token"); err != nil {
			// c.AbortWithStatus(http.StatusUnauthorized)
			c.Redirect(http.StatusPermanentRedirect, "/f/login")
			return
		}

		// Call the next middleware or route handler
		c.Next()
	}
}
