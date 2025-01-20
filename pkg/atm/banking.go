package atm

import "sync"

type Banking struct {
	accounts map[int]*Account
	mu       sync.RWMutex
}

func NewBanking() *Banking {
	return &Banking{
		accounts: make(map[int]*Account),
	}
}

func (b *Banking) CreateAccount(accNo int, balance int) {
	b.accounts[accNo] = NewAccount(accNo, balance)
}

func (b *Banking) GetAccount(accNo int) *Account {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.accounts[accNo]
}

func (b *Banking) RegisterTransaction(transaction Transaction) error {
	return transaction.Execute()
}
