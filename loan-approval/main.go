package main

import (
	"loan-approval/controllers"
	"loan-approval/database"
	_ "loan-approval/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Loan Approval API
// @version 1.0
// @description API for managing and approving loan applications
// @host localhost:8080
// @BasePath /

func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	database.ConnectDatabase()

	r.POST("/apply", controllers.ApplyForLoan)
	r.GET("/applications", controllers.GetApplications)

	r.Run(":8080") // Start server
}
