package model

type User struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Role      string `json:"role" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Age       uint32 `json:"age" validate:"required"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Token     string `json:"token,omitempty"`
}

type UserRequest struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Role      string `json:"role" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Age       uint32 `json:"age" validate:"required"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
type UserProfileRequest struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Age       int    `json:"age" validate:"gte=0"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type JWTResponse struct {
	Token   string `json:"token"`
	Expires string `json:"expired_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
