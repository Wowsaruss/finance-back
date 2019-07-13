package main

import "time"

// Transaction ...
type Transaction struct {
	ID             string    `json:"id"`
	Description    string    `json:"description"`
	Amount         float32   `json:"amount"`
	Date           time.Time `json:"date"`
	Type           string    `json:"type"`
	Monthly        bool      `json:"monthly"`
	Spend          bool      `json:"spend"`
	PaymentType    string    `json:"paymentType"`
	AccountBalance float32   `json:"accountBalance"`
}

// Transactions ...
type Transactions []Transaction
