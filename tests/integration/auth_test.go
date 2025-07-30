package integration_test

import (
	"bytes"
	"encoding/json"
	"gollet/models"
	"gollet/tests"
	"gollet/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginEndpoint(t *testing.T) {
	testDB := tests.SetupTestDB(t)
	defer tests.TeardownTestDB(t, testDB)
	password := "asdasd"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %s\n", err.Error())
	}

	router := tests.GetTestRouter(testDB)

	testUser := models.User{
		Email:        "test@example.com",
		PasswordHash: hashedPassword,
		Role:         models.RoleAdmin,
	}
	testDB.Create(&testUser)

	tests := []struct {
		name         string
		payload      map[string]interface{}
		expectedCode int
	}{
		{
			"Valid login",
			map[string]interface{}{
				"email":    "test@example.com",
				"password": password,
			},
			http.StatusOK,
		},
		{
			"Invalid password",
			map[string]interface{}{
				"email":    "test@example.com",
				"password": "wrongpassword",
			},
			http.StatusUnauthorized,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			payload, _ := json.Marshal(test.payload)
			w := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
			if err != nil {
				t.Fatal("Failed to create new request:", err.Error())
			}

			req.Header.Set("content-type", "application/json")

			router.ServeHTTP(w, req)

			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}
