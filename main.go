package main

import (
	"BasicCrud/controllers"
	"BasicCrud/initilizers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initilizers.LoadEnvVars()
	initilizers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.PostCreate)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/post/:id", controllers.GetSinglePosts)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
