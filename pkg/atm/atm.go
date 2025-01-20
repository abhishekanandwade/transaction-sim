package atm

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

type Atm struct {
	banking  *Banking
	despense *CashDespenser
}

func NewAtm(banking *Banking, despense *CashDespenser) *Atm {
	return &Atm{
		banking:  banking,
		despense: despense,
	}
}

func (a *Atm) Authenticate(card *Card) bool {
	if card.cardNo == 1234 && card.pin == 123 {
		return true
	}
	return false
}

func (a *Atm) WithDrawCash(accountNumber int, amount int) error {
	acc := a.banking.GetAccount(accountNumber)
	if acc == nil {
		return errors.New("account not found")
	}
	if err := a.despense.DespenseCash(amount); err != nil {
		return err
	}
	// if err := a.banking.accounts[acc.accountNumber].Debit(amount); err != nil {
	// 	return err
	// }
	transaction := NewWithdrawTransaction(a.generateTransactionID(), a.banking.accounts[acc.accountNumber], amount)
	a.banking.RegisterTransaction(transaction)

	return nil
}

func (a *Atm) DepositCash(accountNumber int, amount int) error {
	acc := a.banking.GetAccount(accountNumber)
	if acc == nil {
		return errors.New("account not found")
	}
	// if err := a.banking.accounts[acc.accountNumber].Credit(amount); err != nil {
	// 	return err
	// }
	transaction := NewDepositTransaction(a.generateTransactionID(), a.banking.accounts[acc.accountNumber], amount)
	a.banking.RegisterTransaction(transaction)
	return nil
}

func (a *Atm) CheckBalance(accountNumber int) {
	acc := a.banking.GetAccount(accountNumber)
	if acc == nil {
		fmt.Println("account not found")
	}
	fmt.Println(acc.GetAccountBalance())

}

func (a *Atm) generateTransactionID() string {
	value := int64(2)
	txnNumber := atomic.AddInt64(&value, 1)
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("TXN%s%010d", timestamp, txnNumber)
}
