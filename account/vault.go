package account

import (
	"encoding/json"
	"errors"
	"strings"

	"time"

	"github.com/ANU8151/password-manager/files"
	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	db := files.NewJsonDb("accounts.json")
	content, err := db.Read()
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(content, &vault)
	if err != nil {
		color.Red(err.Error())
	}
	return &vault
}

func (vault *Vault) FindAccountByUrl(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *Vault) DeleteAccount(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.save()
	return isDeleted
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (acc *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, errors.New("JSON_MARSHAL_ERROR")
	}
	return file, nil
}

func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red(err.Error())
	}
	db := files.NewJsonDb("accounts.json")
	db.Write(data)
}
