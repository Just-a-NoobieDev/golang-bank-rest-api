package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    int64  `json:"account_number"`
	Balance   int64  `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateAccount struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:   int64(rand.Intn(10000000)),
		Balance: 0,
		CreatedAt: time.Now().UTC(),
	}
}

type Transfer struct {
	ID    int    `json:"id"`
	From   int64 `json:"from"`
	To     int64 `json:"to"`
	Amount int64 `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateTransfer struct {
	From   int64 `json:"from"`
	To     int64 `json:"to"`
	Amount int64 `json:"amount"`
}

func NewTransfer(from, to, amount int64) *Transfer {
	return &Transfer{
		From:   from,
		To:     to,
		Amount: amount,
		CreatedAt: time.Now().UTC(),
	}
}
