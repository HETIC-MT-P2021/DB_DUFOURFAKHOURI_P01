package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SooS just send simple text text on /api route
func SooS(c *gin.Context) {
	c.JSON(http.StatusOK, "SooS")
}
