package handler

import (
	"go-bake/auth"
	"go-bake/helper"
	"go-bake/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func MiddleWare(userService user.Service, authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("authorization")

		if len(authHeader) == 0 || authHeader == "" {
			errorResponse := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "unauthorize user"})

			c.AbortWithStatusJSON(401, errorResponse)
			return
		}

		token, err := authService.ValidateToken(authHeader)

		if err != nil {
			errorResponse := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": err.Error()})

			c.AbortWithStatusJSON(401, errorResponse)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			errorResponse := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "unauthorize user"})

			c.AbortWithStatusJSON(401, errorResponse)
			return
		}

		userID := int(claim["user_id"].(float64))

		c.Set("currentUser", userID)
	}
}
