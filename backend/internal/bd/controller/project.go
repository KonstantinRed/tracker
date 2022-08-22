package controller

import (
	"main_mod/internal/bd/models"
)

func AllProgect() []models.FProject {
	return models.Project{}.All()
}

func GetIndexProject(i int) models.FProject {
	return models.Project{}.GetIndex(i)
}
