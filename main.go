package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var chars = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%&*()")

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimestamp struct {
	account
	createdAt time.Time
	updatedAt time.Time
}

func (acc *accountWithTimestamp) outputPassword() {
	fmt.Println(acc.login)
	fmt.Println(acc.password)
	fmt.Println(acc.url)
}

func (acc *accountWithTimestamp) generatePassword(length int) {
	newPass := make([]rune, length)
	for i := range newPass {
		newPass[i] = chars[rand.IntN(len(chars))]
	}
	acc.password = string(newPass)
}

func newAccountWithTimestamp(login, password, urlString string) (*accountWithTimestamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	fmt.Println("Account created successfully")

	newAcc := &accountWithTimestamp{
		account: account{
			login:    login,
			password: password,
			url:      urlString,
		},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
	newAcc.generatePassword(12)
	return newAcc, nil
}

func main() {
	login := promptData("Enter Login")
	password := promptData("Enter Password")
	url := promptData("Enter URL")

	myAccount, err := newAccountWithTimestamp(login, password, url)
	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}
	myAccount.generatePassword(12)
	myAccount.outputPassword()

}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)

	return res
}
