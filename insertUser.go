package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func insertData(w http.ResponseWriter, r *http.Request) {
	cfg := NewConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Printf("%v", r.Body)
	sqlStatement := `INSERT INTO users (age, email, first_name, last_name)
	VALUES (30, 'hidee@calhoun.io', 'Andrea', 'Hayes')`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Inserted User!!")
	fmt.Println("Inserted successfully!")
}
