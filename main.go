package main

import (
	"fmt"

	"github.com/Sevasit/go-crud/controllers"
	"github.com/Sevasit/go-crud/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()
	r.GET("/employee", controllers.ReadEmployees)
	r.POST("/employee", controllers.CreateEmployee)
	r.GET("/employee/:id", controllers.ReadEmployeesById)
	r.DELETE("/employee/:id", controllers.DeleteEmployeeById)
	r.Run(":5000")
}
