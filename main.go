package main

import (
	"atm-instance/pkg/atm"
)

func main() {
	banking := atm.NewBanking()
	despenser := atm.NewCashDespenser(10000)
	atm := atm.NewAtm(banking, despenser)

	//add new account
	banking.CreateAccount(1233123, 4000)

	atm.CheckBalance(1233123)

	atm.WithDrawCash(1233123, 500)

	atm.CheckBalance(1233123)

	atm.DepositCash(1233123, 2000)

	atm.CheckBalance(1233123)
}
