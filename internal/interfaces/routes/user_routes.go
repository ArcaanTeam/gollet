package routes

import (
	"gollet/internal/entities"
	"gollet/internal/interfaces/controllers"
	"gollet/internal/interfaces/middlewares"
	"gollet/internal/repositories"
	"gollet/internal/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRoutes struct {
	db     *gorm.DB
	engine *gin.Engine
}

func NewUserRoutes(db *gorm.DB, engine *gin.Engine) *UserRoutes {
	return &UserRoutes{
		engine: engine,
		db:     db,
	}
}

func (r *UserRoutes) Setup() {
	userRepo := repositories.NewUserRepository(r.db)
	userService := usecases.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	usersGroup := r.engine.Group("/users")
	usersGroup.Use(middlewares.JwtAuthMiddleware(), middlewares.RBAC(entities.RoleAdmin))
	{
		usersGroup.POST("", userController.CreateUser)
		usersGroup.GET("/:id", userController.GetUserByID)
		usersGroup.PUT("/:id", userController.UpdateUser)
	}

	protectedRoutes := r.engine.Group("/")
	protectedRoutes.Use(middlewares.JwtAuthMiddleware())
	{
		protectedRoutes.GET("/profile", userController.GetProfile)
	}
}
