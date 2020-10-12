package controllers

import (
	"fmt"
	"net/http"

	"github.com/SteakBarbare/DB_DUFOURFAKHOURI_P01/models"
	"github.com/SteakBarbare/DB_DUFOURFAKHOURI_P01/repository"
	"github.com/gin-gonic/gin"
)

func GetOffices(c *gin.Context) {
	var offices []models.Employee
	var err error

	offices, err = repository.FindOffices()

	fmt.Println(err)
	fmt.Println("employee selected: /n", offices)

	if err != nil {
		c.JSON(http.StatusNotFound, "Couldn't fetch offices.")
		return
	}

	c.JSON(http.StatusOK, offices)
}
