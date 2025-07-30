package routes

import (
	"gollet/api/controllers"
	"gollet/repositories"
	"gollet/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthRoutes struct {
	db     *gorm.DB
	engine *gin.Engine
}

func NewAuthRoutes(db *gorm.DB, engine *gin.Engine) *AuthRoutes {
	return &AuthRoutes{
		db:     db,
		engine: engine,
	}
}

func (r *AuthRoutes) Setup() {
	authRepo := repositories.NewAuthRepository(r.db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	r.engine.POST("/login", authController.Login)
}
