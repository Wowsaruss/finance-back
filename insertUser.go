package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Wowsaruss/financial-back-go/pkg/config"
)

func insertData(w http.ResponseWriter, r *http.Request) {
	cfg := config.NewConfig()

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
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `INSERT INTO transactions (date, description, amount, account_balance, type, payment_type, monthly, spend)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err = db.Exec(sqlStatement, transaction.Date, transaction.Description, transaction.Amount, transaction.AccountBalance, transaction.Type, transaction.PaymentType, transaction.Monthly, transaction.Spend)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Inserted User!!")
}
