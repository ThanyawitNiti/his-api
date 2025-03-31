package models

import "time"

type Staff struct {
	IsDelete   bool      `json:"is_delete"`
	IsEdit     bool      `json:"is_edit"`
	CreatedAt  time.Time `json:"created_at"`
	EditAt     time.Time `json:"edit_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	HospitalId int8      `json:"hospital_id"`
}
