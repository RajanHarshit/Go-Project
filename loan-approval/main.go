package main

import (
	"loan-approval/controllers"
	"loan-approval/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()

	r.POST("/apply", controllers.ApplyForLoan)
	r.GET("/applications", controllers.GetApplications)

	r.Run(":8080") // Start server
}
