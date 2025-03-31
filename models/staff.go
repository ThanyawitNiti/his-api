package models

import "time"

type Staff struct {
	IsDelete     bool      `json:"is_delete"`
	IsEdit       bool      `json:"is_edit"`
	CreatedAt    time.Time `json:"created_at"`
	EditAt       time.Time `json:"edit_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	FirstNameTH  string    `json:"first_name_th"`
	MiddleNameTH string    `json:"middle_name_th"`
	LastNameTH   string    `json:"last_name_th"`
	FirstNameEN  string    `json:"first_name_en"`
	MiddleNameEN string    `json:"middle_name_en"`
	LastNameEN   string    `json:"last_name_en"`
	DateOfBirth  string    `json:"date_of_birth"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	Gender       string    `json:"gender"`
	HospitalId   int8      `json:"hospital_id"`
	Department   string    `json:"department"`
	Position     string    `json:"position"`
}
