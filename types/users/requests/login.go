package requests

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
}
