package main

import (
	"BasicCrud/initilizers"
	"BasicCrud/models"
	"log"
)

func init() {
	initilizers.ConnectToDb()
	initilizers.LoadEnvVars()
}

func main() {
	err := initilizers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("unable to initialize the database!")
	}
}
