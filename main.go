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
	authController := controllers.NewAuthController(db.DB)

	r := gin.Default()

	r.POST("/users", userController.CreateUser)
	r.POST("/login", authController.Login)

	r.Run(":8080")
}
