package main

import (
	"fmt"

	"github.com/ANU8151/password-manager/account"
	"github.com/ANU8151/password-manager/files"
)

func main() {
	files.WriteFile("Hello!!!", "account.txt")
	login := promptData("Enter Login")
	password := promptData("Enter Password")
	url := promptData("Enter URL")

	myAccount, err := account.NewAccountWithTimestamp(login, password, url)
	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}
	myAccount.OutputPassword()

}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)

	return res
}
