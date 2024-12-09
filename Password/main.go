package main

import (
	"fmt"
	"math/rand"
)

// struct - это по сути своей класс в других яп
type account struct {
	login    string
	password string
	url      string
}

// Это метод который принадлежит структуре
func (account) outputPassword(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)
}

var letterRunse = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func main() {

	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	//Правильная форма структуры!
	myAccount := account{
		login:    login,
		password: password,
		url:      url,
	}

	outputPassword(&myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)

	return res
}

func generatePassword(n int) string {
	res := make([]rune, n)

	for i := range res {
		res[i] = letterRunse[rand.Intn(len(letterRunse))]
	}

	return string(res)
}
