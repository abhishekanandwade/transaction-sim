package atm

type Transaction interface {
	Execute() error
}

type BaseTransaction struct {
	transactionNo string
	account       *Account
	amount        int
}
