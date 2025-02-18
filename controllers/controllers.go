package controllers

import (
	"BasicCrud/initilizers"
	"BasicCrud/models"
	"log"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	//get data from the req body
	var PostBody struct {
		Title string
		Body  string
	}
	// bind it to the gin server
	c.Bind(&PostBody)

	//create a post
	post := models.Post{Title: PostBody.Title, Body: PostBody.Body}
	result := initilizers.DB.Create(&post)

	if result.Error != nil {
		log.Fatal("error while creating a post!")
		c.Status(400)
		return
	}

	//return it
	c.JSON(200, gin.H{
		"post":    post,
		"message": "successfully created the post",
	})
}

func GetPosts(c *gin.Context) {
	//create a slice of receving struct
	var posts []models.Post
	//Get the posts
	result := initilizers.DB.Find(&posts)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	c.JSON(200, gin.H{
		"Posts": posts,
	})
}

func GetSinglePosts(c *gin.Context) {
	// request data bind
	id := c.Param("id")

	//create a slice of receving struct
	var post models.Post
	result := initilizers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	c.JSON(200, gin.H{
		"Post": post,
	})
}

func UpdatePost(c *gin.Context) {
	//get the id
	id := c.Param("id")

	//get the body of req data
	var PostBody struct {
		Title string
		Body  string
	}

	c.Bind(&PostBody)

	//find the post we were updating
	var post models.Post
	initilizers.DB.First(&post, id)

	//Update the post
	initilizers.DB.Model(&post).Updates(models.Post{
		Title: PostBody.Title,
		Body:  PostBody.Body,
	})

	//Return it
	c.JSON(200, gin.H{
		"Updated Post": post,
	})
}

func DeletePost(c *gin.Context) {
	//get the id
	id := c.Param("id")
	//delete the post
	initilizers.DB.Delete(&models.Post{}, id)
	//respond
	c.JSON(200, gin.H{
		"Deleted Post": id,
	})
}
