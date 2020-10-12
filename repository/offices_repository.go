package repository

import (
	"github.com/SteakBarbare/DB_DUFOURFAKHOURI_P01/database"
	"github.com/SteakBarbare/DB_DUFOURFAKHOURI_P01/models"
	"github.com/SteakBarbare/DB_DUFOURFAKHOURI_P01/utils"
)

func FindOffices() ([]models.Office, error) {

	var (
		OfficeCode                                                                     uint64
		City, Phone, AddressLine1, AddressLine2, State, Country, PostalCode, Territory string
	)

	rows, err := database.DB.Query("SELECT * FROM offices")

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var offices []models.Office

	for rows.Next() {

		err = rows.Scan(
			&OfficeCode,
			&City,
			&Phone,
			&AddressLine1,
			&AddressLine2,
			&State,
			&Country,
			&PostalCode,
			&Territory,
		)

		office := models.Office{
			OfficeCode:   OfficeCode,
			City:         City,
			Phone:        Phone,
			AddressLine1: AddressLine1,
			AddressLine2: AddressLine2,
			State:        State,
			Country:      Country,
			PostalCode:   PostalCode,
			Territory:    Territory,
		}

		offices = append(offices, office)
	}

	if err != nil {
		//panic(err.Error())
	}

	return offices, nil
}

func FindEmployeesByOffice(code uint64) ([]*models.Employee, error) {
	var (
		EmployeeNumber                                              uint64
		Lastname, Firstname, Extension, Email, OfficeCode, JobTitle string
		ReportsTo                                                   utils.NullString
	)

	rows, err := database.DB.Query("SELECT e.employeeNumber, e.firstName, e.lastName, e.extension, e.email, e.officeCode, e.jobTitle, e.reportsTo FROM employees e INNER JOIN offices o ON e.officeCode = o.officeCode WHERE o.officeCode = ?;", code)

	defer rows.Close()

	if err != nil {
		panic(err.Error())
	}

	var employees []*models.Employee

	for rows.Next() {
		err = rows.Scan(
			&EmployeeNumber,
			&Lastname,
			&Firstname,
			&Extension,
			&Email,
			&OfficeCode,
			&JobTitle,
			&ReportsTo)

		employee := &models.Employee{
			EmployeeNumber: EmployeeNumber,
			Lastname:       Lastname,
			Firstname:      Firstname,
			Extension:      Extension,
			Email:          Email,
			OfficeCode:     OfficeCode,
			JobTitle:       JobTitle,
			ReportsTo:      ReportsTo.String,
		}
		employees = append(employees, employee)
	}
	if err != nil {
		//panic(err.Error())
	}

	return employees, nil
}
