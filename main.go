package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

var chars = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%&*()")

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login)
	fmt.Println(acc.password)
	fmt.Println(acc.url)
}

func (acc *account) generatePassword(length int) {
	newPass := make([]rune, length)
	for i := range newPass {
		newPass[i] = chars[rand.IntN(len(chars))]
	}
	acc.password = string(newPass)
}

func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	fmt.Println("Account created successfully")

	newAcc := &account{
		login:    login,
		password: password,
		url:      urlString,
	}
	newAcc.generatePassword(12)
	return newAcc, nil
}

func main() {
	login := promptData("Enter Login")
	password := promptData("Enter Password")
	url := promptData("Enter URL")

	myAccount, err := newAccount(login, password, url)
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
