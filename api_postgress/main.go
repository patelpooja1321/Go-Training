package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pooja"
	dbname   = "first_db"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port =%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	checkError(err)

	defer db.Close()
	insertstmt := `insert into "Employee"("Name","EmpId") values('Rohit',21)`
	_, e := db.Exec(insertstmt)
	checkError(e)

	inserDynstmt := `insert into "Employee"("Name","EmpId") values($1, $2)`
	_, e = db.Exec(inserDynstmt, "Krish", 05)
	checkError(e)

}
func checkError(err error) {
	if err != nil {
		panic(err)
	}

}
