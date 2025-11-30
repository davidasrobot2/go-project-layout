package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseDomain struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (base *BaseDomain) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.New()
	return nil
}

func (base *BaseDomain) BeforeUpdate(tx *gorm.DB) error {
	base.UpdatedAt = new(time.Time)
	*base.UpdatedAt = time.Now()
	return nil
}

type AuthToken struct {
	Token string `json:"token"`
}
