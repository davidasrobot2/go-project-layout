package repository

import (
	"davidasrobot2/go-boilerplate/internal/domain"

	"gorm.io/gorm"
)

type AdministratorRepo struct {
	db *gorm.DB
}

func NewAdministratorRepository(db *gorm.DB) domain.AdministratorRepository {
	return &AdministratorRepo{db: db}
}

func (r *AdministratorRepo) Create(admin *domain.Administrator) error {
	return r.db.Create(admin).Error
}

func (r *AdministratorRepo) FindByEmail(email string) (*domain.Administrator, error) {
	var admin domain.Administrator
	err := r.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AdministratorRepo) FindAll() ([]*domain.Administrator, error) {
	var admins []*domain.Administrator
	err := r.db.Find(&admins).Error
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (r *AdministratorRepo) FindByID(id string) (*domain.Administrator, error) {
	var admin domain.Administrator
	err := r.db.Where("id", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
