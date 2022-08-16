package models

import (
	"context"
	"fmt"
	"log"
	"os"

	pg "github.com/jackc/pgx/v4"
)

type structFields struct {
	Id   int
	Name string
}

type SliceRet []structFields

const (
	connectString = "postgres://postgres:postgres@localhost:5432/Tracker"
)

func Connection() pg.Conn {
	conn, err := pg.Connect(context.Background(), connectString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		log.Panic(err)
	}
	return *conn
}

func CloseConnection(conn *pg.Conn) {
	err := conn.Close(context.Background())

	if err != nil {
		fmt.Fprintf(os.Stderr, "не удалось закрыть конект: %v\n", err)
		log.Panic(err)
	}
}

func Query(conn *pg.Conn, query string, args ...interface{}) pg.Rows {
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Не удалось выполнить запрос: %v\n", err)
		log.Panic(err)
		return nil
	}
	return rows
}

func AllCompany() SliceRet {
	conn := Connection()
	defer CloseConnection(&conn)

	rows := Query(&conn, "select id, name from \"Company\"")

	if rows == nil {
		return SliceRet{}
	}

	var ret []structFields

	for rows.Next() {

		record := structFields{}
		err := rows.Scan(&record.Id, &record.Name)

		if err != nil {
			log.Panic(err)
			return SliceRet{}
		}

		ret = append(ret, record)
		// retByte := json.Marshal(ret)
	}
	return ret
}

func Company(id int) SliceRet {
	conn := Connection()
	defer CloseConnection(&conn)

	query := "select id, name from \"Company\" where id = $1"

	record := structFields{}
	err := conn.QueryRow(context.Background(), query, id).Scan(&record.Id, &record.Name)
	if err != nil {
		log.Panic(err)
	}
	sliceRecord := make(SliceRet, 1)
	sliceRecord[0] = record

	return sliceRecord
}

// func AddCompany(id int) error {
// 	conn := Connection()
// 	defer CloseConnection(&conn)

// 	query := "select id, name from \"Company\" where id = $1"

// 	record := structFields{}
// 	err := conn.QueryRow(context.Background(), query, id).Scan(&record.id, &record.name)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	sliceRecord := make(SliceRet, 1)
// 	sliceRecord[0] = record

// 	return sliceRecord
// }
