package dtos

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
}
