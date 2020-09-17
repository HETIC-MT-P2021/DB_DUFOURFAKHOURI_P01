package router

import (
	"fmt"
	"github.com/HETIC-MT-P2021/DB_DUFOURFAKHOURI_P01/database"
	"github.com/HETIC-MT-P2021/DB_DUFOURFAKHOURI_P01/models"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRequest(){
	router := gin.New()

	bbd := database.DbConn
	repository := models.Repository{Conn: bbd}

	// Global middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/SuuS", func(c *gin.Context){
		fmt.Println("Saucisse ?")
	})

	router.GET("/getCustomers", func(c *gin.Context){
		repository.QueryTest()
	})

	// Listen and serve on 0.0.0.0:8080
	log.Fatal(router.Run(":8080"))
}