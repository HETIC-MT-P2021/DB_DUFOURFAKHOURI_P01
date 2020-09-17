package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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

func dbSetup(){
	db, err := sql.Open("mysql",
		"gobdd:gobdd@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func main() {
	HandleRequest()

	dbSetup()

	//insert, err := db.Query("J'aodre la sacuisse (1, 'SooS')")
}



