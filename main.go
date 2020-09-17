package main

import (
	"github.com/HETIC-MT-P2021/DB_DUFOURFAKHOURI_P01/router"
	"github.com/HETIC-MT-P2021/DB_DUFOURFAKHOURI_P01/database"
)

func main() {
	router.HandleRequest()
	database.DbConnector()
}



