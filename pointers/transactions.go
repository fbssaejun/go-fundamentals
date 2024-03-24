package main

import (
	"errors"
	"fmt"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(customer *customer, transaction *transaction) error {
	switch transactionType := transaction.transactionType; transactionType {
	case transactionDeposit:
		customer.balance += transaction.amount
		return nil
	case transactionWithdrawal:
		if customer.balance < transaction.amount {
			return errors.New("insufficient funds")
		}
		customer.balance -= transaction.amount
		return nil
	default:
		return errors.New("unkown transaction type")
	}
}

func main() {
	customer := customer{
		id:      1,
		balance: 100,
	}

	transaction1 := transaction{
		customerID:      1,
		amount:          50,
		transactionType: transactionWithdrawal,
	}

	transaction2 := transaction{
		customerID:      1,
		amount:          100,
		transactionType: transactionDeposit,
	}

	transaction3 := transaction{
		customerID:      1,
		amount:          100,
		transactionType: "unknown",
	}

	transaction4 := transaction{
		customerID:      1,
		amount:          1000,
		transactionType: transactionWithdrawal,
	}

	updateBalance(&customer, &transaction1)
	fmt.Println(customer.balance)

	updateBalance(&customer, &transaction2)
	fmt.Println(customer.balance)

	fmt.Println(updateBalance(&customer, &transaction3))
	fmt.Println(updateBalance(&customer, &transaction4))

}
