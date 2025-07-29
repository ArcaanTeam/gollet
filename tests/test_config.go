package tests

import (
	"gollet/controllers"
	"gollet/middlewares"
	"gollet/models"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	testDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal("failed to connect to tets database:", err)
	}

	err = testDB.AutoMigrate(&models.User{})

	if err != nil {
		t.Fatal("Failed to migrate test database:", err)
	}

	return testDB
}

func GetTestRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	authController := controllers.NewAuthController(db)
	userController := controllers.NewUserController(db)

	usersGroup := router.Group("/users")
	usersGroup.Use(middlewares.JwtAuthMiddleware(), middlewares.RBAC(models.RoleAdmin))
	{
		usersGroup.POST("", userController.CreateUser)
		usersGroup.GET("", userController.GetUsers)
	}

	router.POST("/login", authController.Login)

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middlewares.JwtAuthMiddleware())
	{
		protectedRoutes.GET("/profile", userController.GetProfile)
	}

	return router
}

func TeardownTestDB(t *testing.T, db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal("Failed to get generic database interface:", err.Error())
	}

	err = sqlDB.Close()
	if err != nil {
		t.Fatal("Failed to close database connection:", err)
	}
}
