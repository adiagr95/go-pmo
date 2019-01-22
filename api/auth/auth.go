package auth

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup){
	auth := r.Group("/auth")
	{
		auth.POST("/login", login)
		auth.POST("/register", register)
	}
}
