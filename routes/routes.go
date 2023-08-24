package routes

func LeadRoutes(app *gin.Engine) {
	lead := app.Group("/lead")

	leadDetail := lead.Group("/:id")
	// Create Lead Endpoint
	lead.POST("", controllers.CreateLead)
	// Get one Lead Detail information Endpoint
	leadDetail.GET("", controllers.GetLead)
	// Delete Lead Endpoint
	leadDetail.DELETE("", controllers.DeleteLead)
	// Update Lead Endpoint
	leadDetail.PUT("", controllers.UpdateLead)

}