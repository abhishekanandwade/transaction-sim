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

func (b *Banking) ChackBalance(accNo int) int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.accounts[accNo].balance
}

func (b *Banking) CreditAmount(accNo int, amt int) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if err := b.accounts[accNo].Credit(amt); err != nil {
		return err
	}

	return nil
}

func (b *Banking) DebitAmount(accNo int, amt int) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if err := b.accounts[accNo].Debit(amt); err != nil {
		return err
	}

	return nil
}
