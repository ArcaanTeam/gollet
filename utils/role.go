package utils

import "gollet/models"

func IsValidRole(role string) bool {
	return role == models.RoleAdmin ||
		role == models.RoleUser
}
