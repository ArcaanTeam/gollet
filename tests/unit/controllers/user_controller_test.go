package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gollet/models"
	"gollet/tests"
	"gollet/utils"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromoteUser(t *testing.T) {
	testDB := tests.SetupTestDB(t)
	defer tests.TeardownTestDB(t, testDB)

	adminPasswordHash, err := utils.HashPassword("admin-password")
	if err != nil {
		t.Fatal("Unable to hash admin password:", err.Error())
	}
	userPasswordHash, err := utils.HashPassword("user-password")
	if err != nil {
		t.Fatal("Unable to hash user password:", err.Error())
	}

	// Create test admin and regular users
	adminUser := models.User{
		Name:         "admin user",
		Email:        "admin@email.com",
		Role:         models.RoleAdmin,
		PasswordHash: adminPasswordHash,
	}
	if err := testDB.Create(&adminUser).Error; err != nil {
		t.Fatal("Failed to create admin user:", err.Error())
	}

	regularUser := models.User{
		Name:         "regular user",
		Email:        "user@email.com",
		Role:         models.RoleUser,
		PasswordHash: userPasswordHash,
	}
	if err := testDB.Create(&regularUser).Error; err != nil {
		t.Fatal("Failed to create regular user:", err.Error())
	}

	adminToken, err := utils.GenerateJwtToken(
		strconv.Itoa(int(adminUser.ID)),
		adminUser.Email,
		adminUser.Role,
	)
	if err != nil {
		t.Fatal("Failed to generate admin token:", err.Error())
	}

	payload := map[string]interface{}{
		"new_role": "admin",
	}

	w := httptest.NewRecorder()
	marshaledPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("Failed to marshal the payload:", err.Error())
	}
	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"/admin/users/%s/promote",
			strconv.FormatUint(uint64(regularUser.ID), 10),
		),
		bytes.NewBuffer(marshaledPayload),
	)
	req.Header.Set("authorization", "Bearer "+adminToken)
	req.Header.Set("content-type", "application/json")

	router := tests.GetTestRouter(testDB)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedUser models.User
	testDB.First(&updatedUser, "id = ?", regularUser.ID)
	assert.Equal(t, models.RoleAdmin, updatedUser.Role)
}
