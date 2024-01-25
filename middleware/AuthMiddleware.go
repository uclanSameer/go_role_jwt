package middleware

import (
	"backend_01/auth"
	"backend_01/ent"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT secret key
var jwtKey = []byte("secret")

// JWT claims struct
type Claims struct {
	UserID uint
	Role   string
	jwt.StandardClaims
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(user *ent.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: uint(user.ID),
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AuthMiddleware is a gin middleware that checks the JWT token in the request header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			return
		}
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Expired token"})
			return
		}
		// Set the user ID and role in the context
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}

// RoleMiddleware is a gin middleware that checks the user's role and permission for a specific action
func RoleMiddleware(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		// Check if the user's role is allowed to perform the action
		if IsAllowed(role.(string), permission) {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You don't have permission to do this"})
		}
	}
}

// IsAllowed checks if a role has a permission
func IsAllowed(role string, permission string) bool {
	for _, p := range auth.Permissions {
		if p.Name == permission {
			for _, r := range p.Roles {
				if r == role {
					return true
				}
			}
		}
	}
	return false
}
