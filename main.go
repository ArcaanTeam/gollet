package main

import (
	"gollet/api/routes"
	"gollet/config"
	"gollet/db"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	db.Migrate()

	engine := gin.Default()

	userRoutes := routes.NewUserRoutes(db.DB, engine)
	userRoutes.Setup()

	authRoutes := routes.NewAuthRoutes(db.DB, engine)
	authRoutes.Setup()

	engine.Run(":8080")
}
