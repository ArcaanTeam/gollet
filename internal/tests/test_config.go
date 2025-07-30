package tests

import (
	"gollet/internal/entities"
	"gollet/internal/interfaces/routes"
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

	err = testDB.AutoMigrate(&entities.User{})

	if err != nil {
		t.Fatal("Failed to migrate test database:", err)
	}

	return testDB
}

func GetTestRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	engine := gin.New()

	userRoutes := routes.NewUserRoutes(db, engine)
	authRoutes := routes.NewAuthRoutes(db, engine)

	userRoutes.Setup()
	authRoutes.Setup()

	return engine
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
