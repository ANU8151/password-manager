package main

import (
	"fmt"

	"github.com/ANU8151/password-manager/account"
	"github.com/ANU8151/password-manager/files"
)

func main() {
	createAccount()
}

func createAccount() {
	login := promptData("Enter Login")
	password := promptData("Enter Password")
	url := promptData("Enter URL")

	myAccount, err := account.NewAccountWithTimestamp(login, password, url)
	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}

	data, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Error converting account to bytes:", err)
		return
	}

	files.WriteFile(data, "accounts.json")

}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)

	return res
}
