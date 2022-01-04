package dto

type UpdateUserRequest struct {
	ID    int64  `json:"id"`
	Name  string `json:"name" binding:"required,min=1"`
	Email string `json:"email" binding:"required,email"`
}
