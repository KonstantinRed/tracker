package models

import "log"

type Project struct{}

type FProject struct {
	Id           int
	Id_company   int
	Name_project string
	Description  string
}

func (table Project) All() []FProject {
	db := Connection()
	defer CloseConnection(db)

	comp := []FProject{}
	err := db.Select(&comp, "Select * from \"Project\"")

	if err != nil {
		log.Panic(err)
	}

	return comp
}

func (table Project) GetIndex(id int) FProject {
	db := Connection()
	defer CloseConnection(db)

	query := "select * from \"Project\" where id = $1"

	task := FProject{}
	err := db.Get(&task, query, id)

	if err != nil {
		log.Panic(err)
	}

	return task
}
