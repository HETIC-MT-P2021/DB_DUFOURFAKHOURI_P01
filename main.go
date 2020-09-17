package main

import (
	"github.com/HETIC-MT-P2021/DB_DUFOURFAKHOURI_P01/database"
	"github.com/HETIC-MT-P2021/DB_DUFOURFAKHOURI_P01/router"
)

func main() {
	database.DbConnector()
	router.HandleRequest()
}



