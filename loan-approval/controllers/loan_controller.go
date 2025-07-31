package controllers

import (
	"loan-approval/database"
	"loan-approval/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ApplyForLoan godoc
// @Summary Apply for a loan
// @Description Submits a loan application and returns approval status
// @Tags loans
// @Accept json
// @Produce json
// @Param application body models.LoanApplication true "Loan Application"
// @Success 200 {object} models.LoanApplication
// @Router /apply [post]

func ApplyForLoan(c *gin.Context) {
	var app models.LoanApplication

	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simple approval logic
	if app.Income >= 30000 && app.CreditScore >= 650 && app.Amount <= app.Income*5 {
		app.Status = "approved"
	} else {
		app.Status = "rejected"
	}

	database.DB.Create(&app)
	c.JSON(http.StatusOK, app)
}

func GetApplications(c *gin.Context) {
	var apps []models.LoanApplication
	database.DB.Find(&apps)
	c.JSON(http.StatusOK, apps)
}
