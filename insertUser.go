package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func insertData(w http.ResponseWriter, r *http.Request) {
	cfg := NewConfig()

	var transaction Transaction

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &transaction); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `INSERT INTO users (id, date, description, amount, account_balance, payment_type, monthly, spend)
	VALUES (transaction.id, transaction.date, transaction.description, transaction.amount, transaction.accountBalance, transaction.type, transaction.monthly, transaction.spend)`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Inserted User!!")
	fmt.Println("Inserted successfully!")
}
