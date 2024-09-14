package controllers

import (
	"errors"
	"net/http"

	"github.com/Sevasit/go-crud/database"
	"github.com/Sevasit/go-crud/model"
	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {
	var employee *model.Employee
	err := c.ShouldBind(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(employee)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"employee": employee,
	})
}

func ReadEmployees(c *gin.Context) {
	var employees []model.Employee
	res := database.DB.Find(&employees)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("authors not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"employees": employees,
	})
}

func ReadEmployeesById(c *gin.Context) {
	id := c.Param("id")
	var employee model.Employee
	res := database.DB.First(&employee, id)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("employee not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"employee": employee,
	})
}

func DeleteEmployeeById(c *gin.Context) {
	id := c.Param("id")
	var employee model.Employee
	res := database.DB.Delete(&employee, id)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("employee not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"employee": employee,
	})
}

func UpdateEmployeeById(c *gin.Context) {
	id := c.Param("id")
	var employee model.Employee
	err := c.ShouldBind(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Model(&employee).Where("id = ?", id).Updates(employee)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("employee not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"employee": employee,
	})
}
