package common

import (
	"os"

	"github.com/coming-chat/go-sui/account"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetAccount() *account.Account {
	words := os.Getenv("words")
	account, err := account.NewAccountWithMnemonic(words)
	PanicIfError(err)
	return account
}
