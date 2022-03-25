package main

import (
	"database/sql"
	"fmt"
	//"github.com/ALPetrov/authorization/request"
	//"github.com/ALPetrov/authorization/funcadd"
	_"github.com/go-sql-driver/mysql"
)

type tableUser struct {
	id int 
	name     string
	lastName string
	login    string
	password string
	deleted  string
}
type query struct {
	name  string
	lastName string
	login    string
	password string
	deleted  string
}

func main() {

	a := query{
	name: "A4",
	lastName: "P4",
	login: "AP4",
	password: "4",
	deleted: "No", 
	}
	
	db, err := sql.Open("mysql", "user:mysql2@/tcp(86.57.217.99:3306)/testbd")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "insert into testbd.table1 (name, lastName,login, password, deleted) values (?, ?, ?, ?, ?)"
	result, err := db.Exec(query, a.name, a.lastName, a.login, a.password, a.deleted)

	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId())

	rows, err := db.Query("SELECT * FROM testbd.table1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := []tableUser{}

	for rows.Next() {
		p := tableUser{}
		
		err := rows.Scan(&p.id, &p.name, &p.lastName, &p.login, &p.password, &p.deleted)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}

	for _, p := range users {

		fmt.Println(p.id, p.name, p.lastName, p.login, p.password, p.deleted)
	}
}
