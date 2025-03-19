package main

import (
	"fmt"
	"math/rand/v2"
)

var chars = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%&*()")

type account struct {
	login    string
	password string
	url      string
}

func main() {
	fmt.Println(generatePassword(12))

	login := promptData("Enter Login")
	password := promptData("Enter Password")
	url := promptData("Enter URL")

	myAccount := account{
		login:    login,
		password: password,
		url:      url,
	}

	outputPassword(&myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)

	return res
}

func outputPassword(account *account) {
	fmt.Println(account.login)
	fmt.Println(account.password)
	fmt.Println(account.url)
}

func generatePassword(length int) string {
	newPass := make([]rune, length)
	for i := range newPass {
		newPass[i] = chars[rand.IntN(len(chars))]
	}
	return string(newPass)
}
