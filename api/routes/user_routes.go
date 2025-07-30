package routes

import (
	"gollet/api/controllers"
	"gollet/api/middlewares"
	"gollet/models"
	"gollet/repositories"
	"gollet/services"

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
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	usersGroup := r.engine.Group("/users")
	usersGroup.Use(middlewares.JwtAuthMiddleware(), middlewares.RBAC(models.RoleAdmin))
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
