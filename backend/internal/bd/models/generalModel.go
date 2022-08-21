package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type methodsControl interface {
	All()
	GetIndex()
}

func Connection() *sqlx.DB {
	BDconnectString := os.Getenv("BDconnectString")
	Driver := os.Getenv("Driver")

	db, err := sqlx.Connect(Driver, BDconnectString)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func CloseConnection(bd *sqlx.DB) {
	err := bd.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "не удалось закрыть конект: %v\n", err)
		log.Panic(err)
	}
}
