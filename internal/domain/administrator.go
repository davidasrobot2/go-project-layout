package domain

import (
	"context"
	"davidasrobot2/go-boilerplate/pkg/constant"
)

type Administrator struct {
	BaseDomain
	Name     string                      `gorm:"type:varchar(100);not null"`
	Email    string                      `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string                      `gorm:"type:varchar(255);not null"`
	Status   bool                        `gorm:"type:boolean;not null"`
	Level    constant.AdministratorLevel `gorm:"type:int(2);not null"`
}

type AdministratorLoginForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AdministratorRepository interface {
	FindByEmail(email string) (*Administrator, error)
	FindAll() ([]*Administrator, error)
	FindByID(id string) (*Administrator, error)
}

type AdministratorUsecase interface {
	SignIn(ctx context.Context, email, password string) (*AuthToken, error)
}
