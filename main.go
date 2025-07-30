package main

import (
	"gollet/api/controllers"
	"gollet/api/middlewares"
	"gollet/config"
	"gollet/db"
	"gollet/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	db.Migrate()

	userController := controllers.NewUserController(db.DB)
	authController := controllers.NewAuthController(db.DB)

	r := gin.Default()

	usersGroup := r.Group("/users")
	usersGroup.Use(middlewares.JwtAuthMiddleware(), middlewares.RBAC(models.RoleAdmin))
	{
		usersGroup.POST("", userController.CreateUser)
		usersGroup.GET("", userController.GetUsers)
		usersGroup.PUT("/:id/promote", userController.PromoteUser)
	}

	r.POST("/login", authController.Login)

	protectedRoutes := r.Group("/")
	protectedRoutes.Use(middlewares.JwtAuthMiddleware())
	{
		protectedRoutes.GET("/profile", userController.GetProfile)
	}

	r.Run(":8080")
}
