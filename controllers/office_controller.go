package controllers

import (
	"net/http"

	"github.com/SteakBarbare/DB_DUFOURFAKHOURI_P01/models"
	"github.com/SteakBarbare/DB_DUFOURFAKHOURI_P01/repository"
	"github.com/SteakBarbare/DB_DUFOURFAKHOURI_P01/utils"
	"github.com/gin-gonic/gin"
)

func GetOfficeByCode(c *gin.Context) {
	var offices *models.Office
	var err error
	id := utils.ParseStringToUint64(c.Param("id"))

	offices, err = repository.FindOfficeByCode(id)

	if err != nil {
		c.JSON(http.StatusNotFound, "Couldn't fetch offices by code.")
		return
	}

	c.JSON(http.StatusOK, offices)
}
