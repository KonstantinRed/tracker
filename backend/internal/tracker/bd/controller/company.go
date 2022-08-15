package controller

import (
	"main_mod/internal/tracker/bd/models"
)

func Company(i int) models.Records {
	return models.Company(i)
}

func AllCompany() models.Records {
	return models.AllCompany()
}
