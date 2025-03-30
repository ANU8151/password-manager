package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/ANU8151/password-manager/output"
)

var chars = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%&*()")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) Output() {
	output.PrintError(acc.Login)
	output.PrintError(acc.Password)
	output.PrintError(acc.Url)
}

func (acc *Account) generatePassword(length int) {
	newPass := make([]rune, length)
	for i := range newPass {
		newPass[i] = chars[rand.IntN(len(chars))]
	}
	acc.Password = string(newPass)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	output.PrintError("ACCOUNT_CREATED_SUCCESSFULLY")
	return newAcc, nil
}
