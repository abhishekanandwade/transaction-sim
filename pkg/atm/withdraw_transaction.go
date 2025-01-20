package atm

type WithdrawTransaction struct {
	BasicTransaction
}

func NewWithdrawTransaction(transactionId int, account *Account, amount int) *WithdrawTransaction {
	return &WithdrawTransaction{
		BasicTransaction: BasicTransaction{
			transactionNo: transactionId,
			account:       account,
			amount:        amount,
		},
	}

}

func (wt *WithdrawTransaction) Execute() error {
	if err := wt.account.Debit(wt.amount); err != nil {
		return err
	}

	return nil
}
