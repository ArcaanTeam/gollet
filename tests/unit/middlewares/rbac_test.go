package middlewares_test

import (
	"gollet/middlewares"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRBACMiddleware(t *testing.T) {
	tests := []struct {
		name          string
		userRole      string
		requiredRoles []string
		expectedCode  int
	}{
		{"Admin access", "admin", []string{"admin"}, http.StatusOK},
		{"User access denied", "user", []string{"admin"}, http.StatusForbidden},
		{"Multiple roles access", "user", []string{"admin", "user"}, http.StatusOK},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			mockHandler := func(c *gin.Context) {
				c.Status(http.StatusOK)
			}

			c.Set("userRole", test.userRole)

			middlewares.RBAC(test.requiredRoles...)(c)

			if !c.IsAborted() {
				mockHandler(c)
			}

			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}
