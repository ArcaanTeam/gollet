package main

import (
	"gollet/api/routes"
	"gollet/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.InitDB()
	config.Migrate()

	engine := gin.Default()

	userRoutes := routes.NewUserRoutes(config.DB, engine)
	userRoutes.Setup()

	authRoutes := routes.NewAuthRoutes(config.DB, engine)
	authRoutes.Setup()

	engine.Run(":8080")
}
