package models

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Address   string             `json:"address" bson:"address"`
	Role      string             `json:"role" bson:"role"`
	Phone     string             `json:"phone" bson:"phone"`
	Age       uint32             `json:"age" bson:"age"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

// LoginInput struct is used for user login input
type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// Validate function to validate LoginInput using go-playground/validator
func (l *LoginInput) Validate() error {
	validate := validator.New()
	return validate.Struct(l)
}

// RegisterInput struct is used for user registration input
type RegisterInput struct {
	Name      string `json:"name" validate:"required,min=3"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Address   string `json:"address" validate:"required"`
	Role      string `json:"role" validate:"required,oneof=instructor student" default:"student"`
	Phone     string `json:"phone" validate:"required"`
	Age       uint32 `json:"age" validate:"required,gt=0"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Validate function to validate RegisterInput using go-playground/validator
func (r *RegisterInput) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
<<<<<<< HEAD

// UserResponse struct is used for user response
type UserResponse struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Role      string `json:"role" validate:"required,oneof=instructor student" `
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UserToken struct {
		Token string `json:"token"`
	}
}

// UserActivityLog struct is used for user activity
type UserActivityLog struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID            string             `json:"user_id" bson:"user_id"`
	CourseID          string             `json:"course_id" bson:"course_id,omitempty"`
	ActivityType      string             `json:"activity_type" bson:"activity_type"`
	ActivityTimestamp string             `json:"activity_timestamp" bson:"activity_timestamp"`
}
=======
>>>>>>> user-service
