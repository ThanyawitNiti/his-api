package models

import "time"

type Patient struct {
	IsDelete     bool      `json:"is_delete"`
	IsEdit       bool      `json:"is_edit"`
	CreatedAt    time.Time `json:"created_at"`
	EditAt       time.Time `json:"edit_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstNameTH  string    `json:"first_name_th"`
	MiddleNameTH string    `json:"middle_name_th"`
	LastNameTH   string    `json:"last_name_th"`
	FirstNameEN  string    `json:"first_name_en"`
	MiddleNameEN string    `json:"middle_name_en"`
	LastNameEN   string    `json:"last_name_en"`
	DateOfBirth  string    `json:"date_of_birth"`
	PatientHN    string    `json:"patient_hn"`
	NationalID   string    `json:"national_id"`
	PassportID   string    `json:"passport_id"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	Gender       string    `json:"gender"`
	HospitalId   int8      `json:"hospital_id"`
}
