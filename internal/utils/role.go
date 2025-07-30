package utils

import "gollet/internal/entities"

func IsValidRole(role string) bool {
	return role == entities.RoleAdmin ||
		role == entities.RoleUser
}
