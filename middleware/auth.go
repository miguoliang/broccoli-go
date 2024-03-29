package middleware

import (
	"fmt"
	"github.com/MicahParks/keyfunc/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
)

// CheckJWT is a gin.HandlerFunc middleware
// that will check the validity of our JWT.
func CheckJWT(role string) gin.HandlerFunc {

	regionID := "us-east-1"
	userPoolID := "us-east-1_wqKTwXm3J"
	jwksURL := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", regionID, userPoolID)

	// Create the keyfunc.Keyfunc.
	jwks, err := keyfunc.NewDefault([]string{jwksURL})
	if err != nil {
		log.Fatalf("Failed to create the keyfunc.Keyfunc: %s", err)
	}

	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				map[string]string{"message": "JWT is missing."},
			)
			return
		}

		// Parse the token.
		token, err := jwt.Parse(tokenString[7:], jwks.Keyfunc)
		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				map[string]string{"message": "JWT is invalid."},
			)
		}

		// Check if the user has the required role.
		claims := token.Claims.(jwt.MapClaims)
		roles := claims["cognito:groups"].([]interface{})
		if roles == nil || !hasRole(roles, role) {
			ctx.AbortWithStatusJSON(
				http.StatusForbidden,
				map[string]string{"message": "JWT does not contain a role."},
			)
		}

		ctx.Next()
	}
}

func hasRole(roles []interface{}, role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}
