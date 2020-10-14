package controllers

import (
	"net/http"

	"github.com/SteakBarbare/DB_DUFOURFAKHOURI_P01/repository"
	"github.com/SteakBarbare/DB_DUFOURFAKHOURI_P01/utils"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/DB_GAUVIN_P01/models"
)

func GetOrderByCode(c *gin.Context) {
	var orders *[]models.OrderProduct
	var err error
	id := utils.ParseStringToUint64(c.Param("id"))

	orders, err = repository.FindOrderByCode(id)

	if err != nil {
		c.JSON(http.StatusNotFound, "Couldn't fetch Order by code.")
		return
	}

	c.JSON(http.StatusOK, orders)
}
