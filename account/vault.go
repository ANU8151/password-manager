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
	Accounts  []account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	content, err := files.ReadFile("accounts.json")
	if err != nil {
		return &Vault{
			Accounts:  []account{},
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

func (vault *Vault) FindAccountByUrl(url string) []account {
	var accounts []account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *Vault) DeleteAccount(url string) {

}

func (vault *Vault) AddAccount(acc account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red(err.Error())
	}
	files.WriteFile(data, "accounts.json")
}

func (acc *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, errors.New("JSON_MARSHAL_ERROR")
	}
	return file, nil
}
