package main

import "fmt"

//struct - это по сути своей класс в других яп
type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	//Правильная форма структуры!
	account1 := account{
		login:    login,
		password: password,
		url:      url,
	}

	outputPassword(account1)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)

	return res
}

func outputPassword(acc account) {
	fmt.Println(acc)
}
