package dtos

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
}
