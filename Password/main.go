package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")
	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)

	if err != nil {
		fmt.Println("Ошибка")
		return
	}

	myAccount.OutputPassword()

	files.ReadFile()
	files.WriteFile()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)

	return res
}
