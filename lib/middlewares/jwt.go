package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	//"../../database/models"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//_ := c.MustGet("db").(*gorm.DB)
		authorization := c.Request.Header.Get("Authorization")
		sp := strings.Split(authorization, "JWT ")
		if len(sp) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		authorization = sp[1]
		secretKey := []byte("your-256-bit-secret")

		jwtToken, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenMap := jwtToken.Claims.(jwt.Claims)
		fmt.Println(tokenMap)

		//var token models.Token
		//if err := db.Where("user_id = ?", token.Claims).First(&user).Error; err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found!"})
		//	return
		//}

	}
}