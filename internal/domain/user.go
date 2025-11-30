package domain

import (
	"context"
)

// User represents a user in the system.
type User struct {
	BaseDomain
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Email    string `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Phone    string `gorm:"type:varchar(20)" json:"phone"`
	Address  string `gorm:"type:text" json:"address"`
	Password string `gorm:"type:varchar(255)" json:"-"` // Omit password from JSON responses
	Status   bool   `gorm:"type:boolean;default:false" json:"status"`
}

type UserCreateForm struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Phone    string `json:"phone" validate:"required,e164,min=13,max=15"`
	Address  string `json:"address" validate:"required"`
}

type UserUpdateForm struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone" validate:"required,e164,min=13,max=15"`
	Address string `json:"address" validate:"required"`
}

type UpdatePasswordForm struct {
	Password string `json:"password" validate:"required,min=8"`
}

type UserSignInForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// UserRepository defines the interface for user data operations.
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindAll(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	FindByIDWithDetail(ctx context.Context, id string) (*User, error)
}

// UserUsecase defines the interface for user business logic.
type UserUsecase interface {
	Create(ctx context.Context, user UserCreateForm) (*User, error)
	SignIn(ctx context.Context, email, password string) (*AuthToken, error)
	FindAll(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, userreq UserUpdateForm, id string) (*User, error)
	Delete(ctx context.Context, id string) error
	ActivateUser(ctx context.Context, id string, password string) (*User, error)
	GetCurrentUser(ctx context.Context, id string) (*User, error)
}
