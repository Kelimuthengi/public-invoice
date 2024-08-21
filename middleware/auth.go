package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/keliMuthengi/invoiving-api/handlers"
)

var (
	claims = jwt.MapClaims{}
)

func AuthenticationMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// get token

		bearer := c.GetHeader("Authorization")
		token := strings.Split(bearer, " ")[1]

		_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, handlers.ResponseHandler{Message: "Invalid token", Status: 1})
			c.Abort()
			return
		}
		c.Next()

	}

}

// // Validate the token and extract the claims
// if claims, ok := token.Claims.(*models.JwtToken); ok && token.Valid {
// 	return claims, nil
// } else {
// 	return nil, fmt.Errorf("invalid token")
// }
