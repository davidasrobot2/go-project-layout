package repository

import "github.com/google/wire"

// Providers is a Wire provider set that provides a new UserRepository.
var Providers = wire.NewSet(NewUserRepository, NewAdministratorRepository)
