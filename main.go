package main

import (
	"./api"
	"./database"
	"github.com/gin-gonic/gin"
)

func main()  {
	db, err := database.Initialize()
	if err != nil {
		panic(err)
	}
	app := gin.Default()
	app.Use(database.Inject(db))
	api.ApplyRoutes(app)
	app.Run(":8000")

}
