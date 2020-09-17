package models

import "fmt"

func (repository *Repository) QueryTest(){
	results := repository.Conn.QueryRow("SELECT customerName FROM customers")

	fmt.Println(results)
}
