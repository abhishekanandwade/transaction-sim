package atm

type WithdrawTransaction struct {
	BaseTransaction
}

func NewWithdrawTransaction(transactionId string, account *Account, amount int) *WithdrawTransaction {
	return &WithdrawTransaction{
		BaseTransaction: BaseTransaction{
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
