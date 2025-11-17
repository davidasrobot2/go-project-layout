package domain

import (
	"context"
	"time"
)

// User represents a user in the system.
type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"-"` // Omit password from JSON responses
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository defines the interface for user data operations.
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
}

// UserUsecase defines the interface for user business logic.
type UserUsecase interface {
	Create(ctx context.Context, name, email, password string) (*User, error)
	Login(ctx context.Context, email, password string) (string, error)
}
