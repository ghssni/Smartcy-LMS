package model

type User struct {
	ID        string      `json:"id,omitempty"`
	Name      string      `json:"name" validate:"required"`
	Email     string      `json:"email" validate:"required"`
	Password  string      `json:"-" validate:"required"`
	Address   string      `json:"address" validate:"required"`
	Role      string      `json:"role" validate:"required"`
	Phone     string      `json:"phone" validate:"required"`
	Age       uint32      `json:"age" validate:"required"`
	CreatedAt string      `json:"created_at,omitempty"`
	UpdatedAt string      `json:"updated_at,omitempty"`
	Token     JWTResponse `json:"credentials,omitempty"`
}

type JWTResponse struct {
	Token   string `json:"token"`
	Expires string `json:"expired_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
