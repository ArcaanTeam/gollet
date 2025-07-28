package main

import (
	"gollet/config"
	"gollet/controllers"
	"gollet/db"
	"gollet/middlewares"

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
	protectedRoutes := r.Group("/")
	protectedRoutes.Use(middlewares.JwtAuthMiddleware())
	{
		protectedRoutes.GET("/profile", userController.GetProfile)
	}

	r.Run(":8080")
}
