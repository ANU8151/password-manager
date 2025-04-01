package main

import (
	"fmt"
	"strings"

	"github.com/ANU8151/password-manager/account"
	"github.com/ANU8151/password-manager/encrypter"
	"github.com/ANU8151/password-manager/files"
	"github.com/ANU8151/password-manager/output"
	"github.com/joho/godotenv"
	// )
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Add Account",
	"2. Find by URL",
	"3. Find by Login",
	"4. Delete Account",
	"5. Exit",
	"===================",
	"Choose variant: ",
}

func main() {

	err := godotenv.Load()
	if err != nil {
		output.PrintError("ENV_FILE_NOT_FOUND")
	}

	vault := account.NewVault(files.NewJsonDb("accounts.vault"), *encrypter.NewEncrypter())
	// vault := account.NewVault(cloud.NewCloudDb("https://files.cloud.com."))
Menu:
	for {
		fmt.Println("===== M E N U =====")

		variant := promptData(menuVariants...)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Enter URL: ")
	accounts := vault.FindAccount(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Enter Login: ")
	accounts := vault.FindAccount(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		output.PrintError("ACCOUNT_NOT_FOUND")
	}
	for _, account := range *accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Enter URL: ")
	isDeleted := vault.DeleteAccount(url)
	if isDeleted {
		output.PrintError("ACCOUNT_SUCCESSFULLY_DELETED")
	} else {
		output.PrintError("ACCOUNT_NOT_FOUND")
	}

}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Enter Login: ")
	password := promptData("Enter Password: ")
	url := promptData("Enter URL: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("ERROR_CREATING_ACCOUNT")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(prompt ...string) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
