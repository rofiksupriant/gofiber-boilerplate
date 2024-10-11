package model

type LoginRequest struct {
	Username string `json:"username" validate:"required,max=50"`
	Password string `json:"password" validate:"required,max=50"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
	Avatar   string `json:"avatar"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,max=50"`
	Password string `json:"password" validate:"required,max=50"`
	Name     string `json:"name" validate:"required,max:100"`
}
