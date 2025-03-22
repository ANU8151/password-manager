package account

import (
	"encoding/json"
	"errors"

	"time"

	"github.com/ANU8151/password-manager/files"
	"github.com/fatih/color"
)

type vault struct {
	Accounts  []account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *vault {
	content, err := files.ReadFile("accounts.json")
	if err != nil {
		return &vault{
			Accounts:  []account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault vault
	err = json.Unmarshal(content, &vault)
	if err != nil {
		color.Red(err.Error())
	}
	return &vault
}

func (vault *vault) AddAccount(acc account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red(err.Error())
	}
	files.WriteFile(data, "accounts.json")
}

func (acc *vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, errors.New("JSON_MARSHAL_ERROR")
	}
	return file, nil
}
