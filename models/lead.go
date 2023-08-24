package models

import "time"

type Lead struct {
	ID             int       `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	FirstName      string    `gorm:"size:256" json:"first_name" csv:"First Name"`
	MiddleName     string    `gorm:"size:256" json:"middle_name" csv:"Middle Name"`
	LastName       string    `gorm:"size:256" json:"last_name" csv:"Last Name"`
	JobTitle       string    `gorm:"size:256" json:"job" csv:"Job Title"`
	Email          string    `gorm:"size:256;not null;index" json:"email" binding:"required,email" csv:"Email"`
	PhoneNumber    string    `json:"phone_number" csv:"Phone Number"`
	City           string    `json:"city" csv:"City"`
	CurrentCompany string    `json:"company" csv:"Current Company"`
	CompanyWebsite string    `json:"company_website" csv:"Company Website"`
	LinkedIn       string    `json:"linkedin" csv:"LinkedIn Profile"`
	Status         string    `json:"status"`
}
