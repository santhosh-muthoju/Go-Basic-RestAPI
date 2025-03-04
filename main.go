package main

import (
	"BasicCrud/controllers"
	"BasicCrud/initilizers"
	"BasicCrud/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initilizers.LoadEnvVars()
	initilizers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/post", middleware.ReqAuth, controllers.PostCreate)
	r.GET("/posts", middleware.ReqAuth, controllers.GetPosts)
	r.GET("/post/:id", middleware.ReqAuth, controllers.GetSinglePosts)
	r.PUT("/post/:id", middleware.ReqAuth, controllers.UpdatePost)
	r.DELETE("/post/:id", middleware.ReqAuth, controllers.DeletePost)
	r.POST("/signup", controllers.UserSignUp)
	r.POST("/login", controllers.UserLogin)
	r.GET("/user", middleware.ReqAuth, controllers.ValidateUser)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
