package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Wowsaruss/financial-back-go/pkg/config"
	"github.com/gorilla/mux"
)

func getType(w http.ResponseWriter, r *http.Request) {
	cfg := config.NewConfig()
	eventID := mux.Vars(r)["type"]

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM transactions WHERE type = $1`

	rows, err := db.Query(sqlStatement, eventID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var t Transaction
	var trans []Transaction
	for rows.Next() {
		var (
			id             int64
			description    string
			date           string
			amount         []uint8
			accountBalance []uint8
			Type           string
			paymentType    string
			monthly        bool
			spend          bool
		)
		if err := rows.Scan(&id, &description, &date, &amount, &accountBalance, &Type, &paymentType, &monthly, &spend); err != nil {
			log.Fatal(err)
		}
		t.Description = description
		// t.Date = date
		// t.Amount = amount
		// t.AccountBalance = accountBalance
		t.Type = Type
		t.PaymentType = paymentType
		t.Monthly = monthly
		t.Spend = spend

		trans = append(trans, t)
	}
	fmt.Printf("%+v", trans)
}
