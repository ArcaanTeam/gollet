package routes

import (
	"gollet/internal/interfaces/controllers"
	"gollet/internal/repositories"
	"gollet/internal/usecases"

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
	authService := usecases.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	r.engine.POST("/login", authController.Login)
}
