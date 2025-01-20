package atm

import (
	"errors"
	"sync"
)

type CashDespenser struct {
	totalCash int
	mu        sync.Mutex
}

func NewCashDespenser(initialCash int) *CashDespenser {
	return &CashDespenser{
		totalCash: initialCash,
	}
}

func (cd *CashDespenser) AddCash(amount int) error {
	cd.mu.Lock()
	defer cd.mu.Unlock()

	cd.totalCash += amount

	return nil
}

func (cd *CashDespenser) DespenseCash(amount int) error {
	cd.mu.Lock()
	defer cd.mu.Unlock()

	if cd.totalCash >= amount {
		cd.totalCash -= amount
		return nil
	}

	return errors.New("Insuffecient amount")
}
