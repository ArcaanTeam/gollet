package payloads

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"omitempty"`
}

type UpdateUserInput struct {
	Name     string `json:"name" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty"`
	Password string `json:"password" binding:"min=6,omitempty"`
	Role     string `json:"role" binding:"omitempty"`
}
