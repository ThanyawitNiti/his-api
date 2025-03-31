package database

import (
	"gin/models"

	"gorm.io/gorm"
)

type Database interface {
	CreateStaff(staff *models.Staff) error
}

type GormDB struct {
	DB *gorm.DB
}

func (g *GormDB) CreateStaff(staff *models.Staff) error {
	return g.DB.Create(staff).Error
}
