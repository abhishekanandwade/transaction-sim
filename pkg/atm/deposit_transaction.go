package atm

type DepositTransaction struct {
	BasicTransaction
}

func NewDepositTransaction(trnsactionNo int, account *Account, amount int) *DepositTransaction {
	return &DepositTransaction{
		BasicTransaction: BasicTransaction{
			transactionNo: trnsactionNo,
			account:       account,
			amount:        amount,
		},
	}
}

func (dt *DepositTransaction) Execute() error {
	if err := dt.account.Credit(dt.amount); err != nil {
		return err
	}
	return nil
}
