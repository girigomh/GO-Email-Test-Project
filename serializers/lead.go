package serializers


type LeadSerializer struct {
	FirstName      string    `json:"first_name" binding:"required"`
	MiddleName     string    `json:"middle_name"`
	LastName       string    `json:"last_name" binding:"required"`
	JobTitle       string    `json:"job"`
	Email          string    `json:"email" binding:"required, email"`
	PhoneNumber    string    `json:"phone_number"`
	City           string    `json:"city"`
	CurrentCompany string    `json:"company"`
	CompanyWebsite string    `json:"company_website"`
	LinkedIn       string    `json:"linkedin"`
	Status         string    `json:"status"`
}
