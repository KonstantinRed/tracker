package models

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Task struct{}

type FTask struct {
	Id               int `json:"id"`
	Id_company       int `json:"id_company"`
	Id_project       int `json:"id_project"`
	Name_task        string
	Description_task string
	Hour             int `json:"hour"`
	Minute           int `json:"minute"`
	Link_our_tracker string
	Closed           bool `json:"processed"`
	Total_minutes    int  `json:"total_minutes"`
}

type FTask_View struct {
	Id               int    `json:"id"`
	Company          string `json:"Company"`
	Project          string `json:"Project"`
	Name_task        string `json:"Name_task"`
	Description_task string `json:"Description_task"`
	Hour             int    `json:"Hour"`
	Minute           int    `json:"Minute"`
	Closed           bool   `json:"Closed"`
}

func (table Task) All() []FTask {
	db := Connection()
	defer CloseConnection(db)

	comp := []FTask{}
	err := db.Select(&comp, "Select * from \"Task\"")

	if err != nil {
		log.Panic(err)
	}

	return comp
}

func (table Task) AllView() []FTask_View {
	task_view_sql :=
		`SELECT 
			Task.id as id,
			coalesce(Company.Name, '---') as Company,
			coalesce(Project.name_project, '---') as Project,
			Task.name_task as Name_task,
			Task.Description_task as Description_task,
			Task.Hour as Hour,
			Task.Minute as Minute,
			Task.Closed as Closed
		FROM public."Task" as Task 
			Left Join "Company" as Company 
				ON Company.id = Task.id_company
			LEFT join "Project" as Project
				ON Project.id = Task.id_project
		ORDER BY id`

	db := Connection()
	defer CloseConnection(db)

	comp := []FTask_View{}
	err := db.Select(&comp, task_view_sql)

	if err != nil {
		log.Panic(err)
	}

	return comp
}

func (table Task) GetIndex(id int) FTask {
	db := Connection()
	defer CloseConnection(db)

	query := "select * from \"Task\" where id = $1"

	task := FTask{}
	err := db.Get(&task, query, id)

	if err != nil {
		log.Panic(err)
	}

	return task
}

func (table Task) AddTask(task FTask) error {

	db := Connection()
	defer CloseConnection(db)

	_, err := db.NamedExec(
		`INSERT INTO public."Task"(
			id_company, id_project, name_task, description_task, hour, minute, total_minutes, link_our_tracker, closed)
			VALUES (:id_company, :id_project, :name_task, :description_task, :hour, :minute, :total_minutes, :link_our_tracker, :closed)`, task)

	return err
}

func (table Task) DeleteIndex(id []int) error {

	db := Connection()
	defer CloseConnection(db)

	query, args, err := sqlx.In(`DELETE FROM public."Task" WHERE id in (?);`, id)
	query = db.Rebind(query)
	_, err = db.Query(query, args...)

	return err
}
