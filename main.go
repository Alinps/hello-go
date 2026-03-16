package main

import (
	"fmt"
)

type BankAccount struct {
	owner   string
	balance float64
}

func (b *BankAccount) withdraw(amount float64) bool {
	if amount > b.balance {
		return false
	}
	b.balance = b.balance - amount
	return true
}

func (b *BankAccount) deposit(amount float64) bool {
	if amount < 0 {
		return false
	}
	b.balance = b.balance + amount
	return true
}

func main() {
	account := BankAccount{
		owner:   "Alin",
		balance: 25000,
	}
	withdrawSuccess := account.withdraw(5000)
	if withdrawSuccess {
		fmt.Println("Withdrawn successfully")
		fmt.Println("Current balance: ", account.balance)
	} else {
		fmt.Println("Insufficient balance")
	}
	depositSuccess := account.deposit(50000)
	if depositSuccess {
		fmt.Println("Deposited successfully")
		fmt.Println("Current balance: ", account.balance)
	} else {
		fmt.Println("Please deposite amount larger than 0")
	}

}
