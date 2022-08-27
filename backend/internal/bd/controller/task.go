package controller

import (
	"fmt"
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

func AddTask(task models.FTask) error {
	err := models.Task{}.AddTask(task)
	if err != nil {
		fmt.Errorf("%#v", err)
	}
	return err
}

func DeleteTask(id []int) error {
	err := models.Task{}.DeleteIndex(id)
	if err != nil {
		fmt.Errorf("%#v", err)
	}
	return err
}
