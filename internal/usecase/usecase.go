package usecase

import "github.com/google/wire"

// Providers is a Wire provider set that provides a new UserUsecase.
var Providers = wire.NewSet(NewUserUsecase, NewAdministratorUsecase)
