package finance

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Wowsaruss/financial-back-go/pkg/config"
)

func float64frombytes(bytes []byte) float64 {
	s := string(bytes)
	noMoney := strings.Replace(s, "$", "", -1)
	noComma := strings.Replace(noMoney, ",", "", -1)
	amt, err := strconv.ParseFloat(noComma, 32)

	if err != nil {
		log.Println(err)
	}
	return amt
}

func toBool(s string) bool {
	res, _ := strconv.ParseBool(s)
	return res
}

func filterTransactions(w http.ResponseWriter, r *http.Request) {
	cfg := config.NewConfig()
	tp, _ := r.URL.Query()["type"]
	monthly, _ := r.URL.Query()["monthly"]
	spend, _ := r.URL.Query()["spend"]

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM transactions`
	var mb bool
	var sb bool
	var rows *sql.Rows
	if len(tp) > 0 {
		sqlStatement = `SELECT * FROM transactions WHERE type = $1`
		rows, err = db.Query(sqlStatement, tp[0])
		if err != nil {
			panic(err)
		}
	}
	if len(monthly) > 0 {
		mb = toBool(monthly[0])
		sqlStatement = `SELECT * FROM transactions WHERE type = $1 AND monthly = $2`
		rows, err = db.Query(sqlStatement, tp[0], mb)
		if err != nil {
			panic(err)
		}
	}
	if len(spend) > 0 {
		sb = toBool(spend[0])
		sqlStatement = `SELECT * FROM transactions WHERE type = $1 AND monthly = $2 AND spend = $3`
		rows, err = db.Query(sqlStatement, tp[0], mb, sb)
		if err != nil {
			panic(err)
		}
	}
	if len(tp) == 0 && len(monthly) == 0 && len(spend) == 0 {
		rows, err = db.Query(sqlStatement)
		if err != nil {
			panic(err)
		}
	}
	defer rows.Close()

	var t Transaction
	for rows.Next() {
		var (
			id             int64
			date           string
			description    string
			amount         []uint8
			accountBalance []uint8
			Type           string
			paymentType    string
			monthly        bool
			spend          bool
		)
		if err := rows.Scan(&id, &date, &description, &amount, &accountBalance, &Type, &paymentType, &monthly, &spend); err != nil {
			log.Fatal(err)
		}

		dt, err := time.Parse("0001-01-01T00:00:00Z", date)
		if err != nil {
			log.Println(err)
		}

		t.Description = description
		t.Date = dt
		t.Amount = float64frombytes(amount)
		t.AccountBalance = float64frombytes(accountBalance)
		t.Type = Type
		t.PaymentType = paymentType
		t.Monthly = monthly
		t.Spend = spend

		r, err := json.Marshal(t)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(r)
	}
}
