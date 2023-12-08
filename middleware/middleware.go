package middleware

import (
	"go-bank/token"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")
		c.Next()
		latency := time.Since(t)
		log.Print(latency)
		status := c.Writer.Status()
		log.Println(status)
	}
}

func AuthCheck(userName string, id string) (string, error) {
	log.Println("accounts authcheck", userName)
	token, err := token.GenerateToken(id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
