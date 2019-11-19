package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Wowsaruss/financial-back-go/pkg/config"
	_ "github.com/lib/pq"
)

// Transaction ...
type Transaction struct {
	Details           string
	PostingDate       time.Time
	Description       string
	Amount            float64
	PmtType           string
	Balance           float64
	CheckOrSlipNumber string
	Type              string
	Monthly           bool
	Spend             bool
}

func monthlyBool(m string) bool {
	if m == "TRUE" {
		return true
	}
	return false
}

func spendBool(p float64) bool {
	if p >= 0 {
		return false
	}
	return true
}

func main() {
	arg := os.Args[1]
	cfg := config.NewConfig()

	csvFile, err := os.Open(arg)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var tran Transaction
	var transactions []Transaction

	withoutHeader := csvData[1:]
	for _, each := range withoutHeader {
		t, _ := time.Parse("01/02/06", each[1])
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		amount, err := strconv.ParseFloat(each[3], 32)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		balance, err := strconv.ParseFloat(each[5], 32)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		tran.Details = each[0]
		tran.PostingDate = t
		tran.Description = each[2]
		tran.Amount = amount
		tran.PmtType = each[4]
		tran.Balance = balance
		tran.CheckOrSlipNumber = each[6]
		tran.Type = each[7]
		tran.Monthly = monthlyBool(each[8])
		tran.Spend = spendBool(amount)

		transactions = append(transactions, tran)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()

	count := 0
	for _, item := range transactions {
		sqlStatement := `INSERT INTO transactions (date, description, amount, account_balance, type, payment_type, monthly, spend)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

		_, err = db.Exec(sqlStatement, item.PostingDate, item.Description, item.Amount, item.Balance, item.Type, item.PmtType, item.Monthly, item.Spend)
		if err != nil {
			panic(err)
		}
		count++
	}
	fmt.Printf("%v transactions inserted to db.", count)
}
