package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var chars = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%&*()")

type account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *account) OutputPassword() {
	color.Green(acc.Login)
	color.Green(acc.Password)
	color.Green(acc.Url)
}

func (acc *account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, errors.New("JSON_MARSHAL_ERROR")
	}
	return file, nil
}

func (acc *account) generatePassword(length int) {
	newPass := make([]rune, length)
	for i := range newPass {
		newPass[i] = chars[rand.IntN(len(chars))]
	}
	acc.Password = string(newPass)
}

//func

func NewAccountWithTimestamp(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	fmt.Println("Account created successfully")

	newAcc := &account{

		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
