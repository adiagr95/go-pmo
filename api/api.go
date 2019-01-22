package api

import (
	"./auth"
	"./client"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth.ApplyRoutes(api)
		client.ApplyRoutes(api, )
	}
}