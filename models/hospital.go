package models

import "time"

type Hospital struct {
	IsDelete    bool      `json:"is_delete"`
	IsEdit      bool      `json:"is_edit"`
	CreatedAt   time.Time `json:"created_at"`
	EditAt      time.Time `json:"edit_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	NameTh      string    `json:"name_th"`
	NameEn      string    `json:"name_en"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
}
