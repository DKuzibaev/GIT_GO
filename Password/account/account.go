package account

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/fatih/color"
)

// Global variable
var letterRunse = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

// struct - это по сути своей класс в других яп
type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"URL"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Так выгдядит наследование в Go
// type AccountWithTimeStamp struct {
// 	createdAt time.Time
// 	updatedAt time.Time
// 	Account
// }

// Это метод который принадлежит структуре account
func (acc *Account) Output() {
	color.Green("Ваш логин: %s", acc.Login)
	color.Green("Ваш пароль: %s", acc.Password)
	color.Green("Ваш URL: %s", acc.Url)
	fmt.Println("")
}

// Это метод создает новый пароль
func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)

	for i := range res {
		res[i] = letterRunse[rand.Intn(len(letterRunse))]
	}
	acc.Password = string(res)
}

// Конструктор функции
func NewAccount(log, pas, urlStr string) (*Account, error) {
	if log == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlStr)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	//Создание новой структуры
	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     log,
		Password:  pas,
		Url:       urlStr,
	}

	if pas == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}
