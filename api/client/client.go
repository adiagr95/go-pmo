package client

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../../lib/middlewares"
)

func ApplyRoutes(r *gin.RouterGroup){
	auth := r.Group("/", middlewares.JWTMiddleware())
	{
		auth.POST("/client", client)
	}
}

func client(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"name" : "Pinkblue"})
}
