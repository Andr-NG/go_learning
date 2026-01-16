package main

import (
	"errors"
)

type Account struct {
	Owner   string
	Balance int
}

// Deposit method having a pointer receive to modify original struct data
func (a *Account) Deposit(amount int) error {
	if amount <= 0 {
		return errors.New("Deposited amount cannot be negative or equal to 0")
	}

	a.Balance += amount
	return nil
}

// Withdraw method
func (a *Account) Withdraw(amount int) error {

	if amount > a.Balance {
		return errors.New("insufficient funds")
	}

	if amount <= 0 {
		return errors.New("Amount cannot be negative or equal to 0")
	}

	a.Balance -= amount
	return nil
}
