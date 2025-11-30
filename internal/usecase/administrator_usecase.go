package usecase

import (
	"context"
	"davidasrobot2/go-boilerplate/internal/domain"
	"davidasrobot2/go-boilerplate/pkg/auth"
	"davidasrobot2/go-boilerplate/pkg/constant"
	"davidasrobot2/go-boilerplate/pkg/helper"
)

type administratorUsecase struct {
	administratorRepo domain.AdministratorRepository
	jwtGenerator      *auth.JWTGenerator
}

// NewAdministratorUsecase
func NewAdministratorUsecase(administratorRepo domain.AdministratorRepository, jwtGenerator *auth.JWTGenerator) domain.AdministratorUsecase {
	return &administratorUsecase{
		administratorRepo: administratorRepo,
		jwtGenerator:      jwtGenerator,
	}
}

func (uc *administratorUsecase) SignIn(ctx context.Context, email, password string) (*domain.AuthToken, error) {
	admin, err := uc.administratorRepo.FindByEmail(email)
	if err != nil {
		return nil, constant.ErrorMessageInvalidCredential
	}

	if !helper.CheckPasswordHash(password, admin.Password) {
		return nil, constant.ErrorMessageNotFound
	}

	token, err := uc.jwtGenerator.GenerateToken(admin.ID.String())
	if err != nil {
		return nil, constant.ErrorMessageInternalServerError
	}

	return token, nil
}
