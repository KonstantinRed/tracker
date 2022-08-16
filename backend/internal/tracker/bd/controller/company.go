package controller

import (
	"main_mod/internal/tracker/bd/models"
)

func AllCompany() models.SliceRet {
	return models.AllCompany()
}

func Company(i int) models.SliceRet {
	return models.Company(i)
}
