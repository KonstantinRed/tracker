package controller

import (
	"main_mod/internal/bd/models"
)

func AllTask() []models.FTask {
	return models.Task{}.All()
}

func AllTask_view() []models.FTask_View {
	return models.Task{}.AllView()
}

func GetIndexTask(i int) models.FTask {
	return models.Task{}.GetIndex(i)
}

func AddTask() {

}
