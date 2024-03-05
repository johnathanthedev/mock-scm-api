package dtos

type CreateUserRequest struct {
	// TODO: change to max 50
	Username string `json:"username" validate:"required,min=3,max=20"`
}
