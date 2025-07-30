package main

import (
	"gollet/internal/infra"
	"gollet/internal/interfaces/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.LoadConfig()
	infra.InitDB()
	infra.Migrate()

	engine := gin.Default()

	userRoutes := routes.NewUserRoutes(infra.DB, engine)
	userRoutes.Setup()

	authRoutes := routes.NewAuthRoutes(infra.DB, engine)
	authRoutes.Setup()

	engine.Run(":8080")
}
