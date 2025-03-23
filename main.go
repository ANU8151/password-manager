package main

import (
	"fmt"

	"github.com/ANU8151/password-manager/account"
	"github.com/ANU8151/password-manager/files"
	"github.com/ANU8151/password-manager/output"
	// )
)

func main() {
	vault := account.NewVault(files.NewJsonDb("accounts.json"))
	// vault := account.NewVault(cloud.NewCloudDb("https://files.cloud.com."))
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
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

func findAccount(vault *account.VaultWithDb) {
	url := promptData("Enter URL")
	accounts := vault.FindAccountByUrl(url)
	if len(accounts) == 0 {
		output.PrintError("ACCOUNT_NOT_FOUND")
	}
	for _, account := range accounts {
		account.OutputPassword()
	}

}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Enter URL")
	isDeleted := vault.DeleteAccount(url)
	if isDeleted {
		output.PrintError("ACCOUNT_SUCCESSFULLY_DELETED")
	} else {
		output.PrintError("ACCOUNT_NOT_FOUND")
	}

}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Enter Login")
	password := promptData("Enter Password")
	url := promptData("Enter URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("ERROR_CREATING_ACCOUNT")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)

	return res
}
