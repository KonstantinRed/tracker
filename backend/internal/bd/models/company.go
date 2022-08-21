package models

import (
	"log"
)

type Company struct{}

type FCompany struct {
	Id   int
	Name string
}

// type RecordsCompany []fieldsCompany

func (table Company) All() []FCompany {
	db := Connection()
	defer CloseConnection(db)

	comp := []FCompany{}
	err := db.Select(&comp, "Select * from \"Company\"")

	if err != nil {
		log.Panic(err)
	}

	return comp
}

func (table Company) GetIndex(id int) FCompany {
	db := Connection()
	defer CloseConnection(db)

	query := "select * from \"Company\" where id = $1"

	comp := FCompany{}
	err := db.Get(&comp, query, id)

	if err != nil {
		log.Panic(err)
	}

	return comp
}
