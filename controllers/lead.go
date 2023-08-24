package controllers

func CreateLead(ctx *gin.Context) {
	var leadBody serializers.LeadSerializer
	if err := ctx.ShouldBindJSON(&leadBody); err != nil {
		ctx.JSON(
			http.StatusBadRequest, 
			gin.H{"error": err.Error()}
		)
		return
	}

	campaign := models.Lead{
		FirstName: leadBody.FirstName, 
		MiddleName: leadBody.MiddleName, 
		LastName: leadBody.LastName,
		JobTitle: leadBody.JobTitle,
		Email: leadBody.Email,
		PhoneNumber: leadBody.PhoneNumber,
		City: leadBody.City,
		CurrentCompany: leadBody.CurrentCompany,
		CompanyWebsite: leadBody.CompanyWebsite,
		LinkedIn: leadBody.LinkedIn,
		Status: leadBody.Status,
	}

	if err := initializers.DB.Create(&lead).Error; err != nil {
		ctx.JSON(
			http.StatusInternalServerError, 
			gin.H{"error": err.Error()}
		)
		return
	}
	ctx.JSON(http.StatusOK, lead)
}

func GetLead(ctx *gin.Context) {
	leadID = ctx.Param("id")
	
}

func DeleteLead(ctx *gin.Context) {

}

func UpdateLead(ctx *gin.Context) {

}