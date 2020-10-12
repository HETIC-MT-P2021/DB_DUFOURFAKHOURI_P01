package controllers

import (
	"fmt"
	"net/http"

	"github.com/SteakBarbare/DB_DUFOURFAKHOURI_P01/utils"

	"github.com/SteakBarbare/DB_DUFOURFAKHOURI_P01/models"
	"github.com/SteakBarbare/DB_DUFOURFAKHOURI_P01/repository"
	"github.com/gin-gonic/gin"
)

func GetEmployees(c *gin.Context) {
	var employees []models.Employee
	var err error

	employees, err = repository.FindEmployees()

	fmt.Println(err)
	fmt.Println("employee selected: /n", employees)

	if err != nil {
		c.JSON(http.StatusNotFound, "Couldn't fetch employees.")
		return
	}

	c.JSON(http.StatusOK, employees)
}

func GetEmployeesByOfficeCode(c *gin.Context) {
	var employees []*models.Employee
	var err error
	id := utils.ParseStringToUint64(c.Param("id"))

	employees, err = repository.FindEmployeesByOffice(id)

	if err != nil {
		c.JSON(http.StatusNotFound, "Couldn't fetch employees by office.")
		return
	}

	c.JSON(http.StatusOK, employees)
}
