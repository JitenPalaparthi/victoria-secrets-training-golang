package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	a1 := &Account{AccountNo: "312312", AccountName: "Jiten", Balance: 2300.45}
	if b, err := a1.Debit(2200.1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Balance:", b)
	}

	if b, err := a1.Debit(101.); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Balance:", b)
	}

}

func GenError() error {
	return errors.New("Demo error")
}

type Account struct {
	AccountNo   string
	AccountName string
	Balance     float32
}

func (a *Account) Debit(amount float32) (float32, error) {
	if amount <= a.Balance {
		b := a.Balance - amount
		// fmt.Println(b)
		a.Balance = float32(math.Round(float64(b)))
		return a.Balance, nil

	}

	err := &BankingError{Code: "4004", Message: "Insufficient balance;Current Balance is " + fmt.Sprint(a.Balance)}
	//return 0, errors.New("Insufficient balance")
	//return 0, fmt.Errorf("Insufficient balance")
	//return 0, &BankingError{Code: "4004", Message: "Insufficient balance"}
	return 0, err

}

type BankingError struct {
	Code    string
	Message string
}

func (b *BankingError) Error() string {
	return fmt.Sprintf("Code:%s Message:%s", b.Code, b.Message)
}
