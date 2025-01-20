package atm

import (
	"errors"
	"sync"
)

type Account struct {
	accountNumber int
	balance       int
	mu            sync.Mutex
}

func NewAccount(accNumber int, balance int) *Account {
	return &Account{
		accountNumber: accNumber,
		balance:       balance,
	}
}

func (a *Account) GetAccountNumber() int {
	return a.accountNumber
}

func (a *Account) GetAccountBalance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func (a *Account) Credit(amt int) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.balance += amt

	return nil
}

func (a *Account) Debit(amt int) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.balance >= amt {
		a.balance -= amt
		return nil
	}

	return errors.New("Insuffecient Account Balance")
}
