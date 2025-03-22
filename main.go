package main

import (
	"fmt"

	"github.com/ANU8151/password-manager/account"
)

func main() {
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("===== M E N U =====")
	fmt.Println("1. Add Account")
	fmt.Println("2. Find Accounts")
	fmt.Println("3. Delete Account")
	fmt.Println("4. Exit")
	fmt.Println("====================")
	fmt.Println("Choose variant")
	fmt.Scanln(&variant)
	return variant
}

func findAccount() {
	fmt.Println("Find Account")
}

func deleteAccount() {
	fmt.Println("Delete Account")
}

func createAccount() {
	login := promptData("Enter Login")
	password := promptData("Enter Password")
	url := promptData("Enter URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}

	vault := account.NewVault()
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)

	return res
}
