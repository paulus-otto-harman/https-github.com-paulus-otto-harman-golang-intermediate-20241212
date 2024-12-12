package middleware

import (
	"20241212/class/2/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (m *Middleware) OnlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		isValid, data := validateToken(token, m.secretKey)

		if !isValid {
			handler.BadResponse(c, "Unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		userData := strings.Split(data, ":")
		if userData[1] != "admin" {
			handler.BadResponse(c, "Forbidden", http.StatusForbidden)
			c.Abort()
			return
		}

		c.Next()
	}
}
