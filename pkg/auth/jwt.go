package auth

import (
	"davidasrobot2/go-boilerplate/config"
	"davidasrobot2/go-boilerplate/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/wire"
)

// AuthSet is a Wire provider set that provides a new JWTGenerator.
var Providers = wire.NewSet(NewJWTGenerator)

// JWTGenerator handles the creation of JWTs.
type JWTGenerator struct {
	cfg config.JWTConfig
}

// NewJWTGenerator creates a new JWTGenerator.
func NewJWTGenerator(cfg *config.Config) *JWTGenerator {
	return &JWTGenerator{cfg: cfg.JWT}
}

// GenerateToken creates a new JWT for a given user ID.
func (j *JWTGenerator) GenerateToken(userID string) (*domain.AuthToken, error) {
	var AuthToken domain.AuthToken
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	AuthToken.Token, _ = token.SignedString([]byte(j.cfg.Secret))

	return &AuthToken, nil
}
