package utils_test

import (
	"gollet/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordHashing(t *testing.T) {
	password := "securepassword123"

	t.Run("Hash and verify password", func(t *testing.T) {
		hash, err := utils.HashPassword(password)
		assert.NoError(t, err)
		assert.NotEmpty(t, hash)
		assert.True(t, utils.CheckPasswordHash(password, hash) == nil)
	})

	t.Run("Wrong password verification", func(t *testing.T) {
		hash, _ := utils.HashPassword(password)
		assert.NotNil(t, utils.CheckPasswordHash("wrongpassword", hash))
	})

	t.Run("Empty password", func(t *testing.T) {
		_, err := utils.HashPassword("")
		assert.Error(t, err)
	})
}
