package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/HETIC-MT-P2021/DB_DUFOURFAKHOURI_P01/database"
)

func HandleRequest(){
	router := gin.New()

	// Global middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/SuuS", func(c *gin.Context){
		fmt.Println("Saucisse ?")
	})

	router.GET("/", func(c *gin.Context){
		fmt.Println("WoW")
	})

	// Listen and serve on 0.0.0.0:8080
	log.Fatal(router.Run(":8080"))
}


func main() {
	HandleRequest()
	database.DbConnector()
}



