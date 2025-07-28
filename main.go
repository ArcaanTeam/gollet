package main

import (
	"gollet/config"
	"gollet/controllers"
	"gollet/db"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	db.Migrate()

	userController := controllers.NewUserController(db.DB)

	r := gin.Default()

	r.POST("/users", userController.CreateUser)

	r.Run(":8080")
}
