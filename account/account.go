package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var chars = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%&*()")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) OutputPassword() {
	color.Green(acc.Login)
	color.Green(acc.Password)
	color.Green(acc.Url)
}

func (acc *Account) generatePassword(length int) {
	newPass := make([]rune, length)
	for i := range newPass {
		newPass[i] = chars[rand.IntN(len(chars))]
	}
	acc.Password = string(newPass)
}

//func

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	color.Green("ACCOUNT_CREATED_SUCCESSFULLY")

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
	return newAcc, nil
}
