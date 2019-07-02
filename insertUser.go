package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Wowsaruss/financial-back-go/pkg/config"
)

func insertData(w http.ResponseWriter, r *http.Request) {
	cfg := config.NewConfig()
	fmt.Println(cfg)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

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
