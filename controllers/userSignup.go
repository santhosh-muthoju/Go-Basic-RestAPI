package controllers

import (
	"BasicCrud/initilizers"
	"BasicCrud/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserSignUp(c *gin.Context) {
	//body data
	var body struct {
		Email    string
		Password string
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to read the request body",
		})
		return
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to hash the password",
		})
		return
	}

	//create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initilizers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create the user..",
		})
		return
	}

	//respond
	c.JSON(http.StatusOK, gin.H{
		"msg": "successfully created the user...",
	})

}
