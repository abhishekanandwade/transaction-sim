package atm

type DepositTransaction struct {
	BaseTransaction
}

func NewDepositTransaction(trnsactionNo string, account *Account, amount int) *DepositTransaction {
	return &DepositTransaction{
		BaseTransaction: BaseTransaction{
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
