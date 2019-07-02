package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func insertData(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStatement := `INSERT INTO users (age, email, first_name, last_name)
	VALUES (30, 'jon@calhoun.io', 'Jonathan', 'Calhoun')`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Inserted User!!")
	fmt.Println("Inserted successfully!")
}
