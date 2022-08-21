package controller

import (
	"main_mod/internal/bd/models"
)

func AllCompany() []models.FCompany {
	return models.Company{}.All()
}

func Company(i int) models.FCompany {
	return models.Company{}.GetIndex(i)
}
