package atm

type Transaction interface {
	Execute(amount int) error
}

type BaseTransaction struct {
	transactionNo int
	account       *Account
	amount        int
}
