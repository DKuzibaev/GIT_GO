package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	fmt.Println("___Менеджер паролей___\n")

Menu:
	for {
		variant := getMenu()

		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			removeAccount()
		default:
			break Menu
		}
	}

}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scanln(&res)

	return res
}

func createAccount() {
	login := promptData("Введите логин:")
	password := promptData("Введите пароль:")
	url := promptData("Введите URL:")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		fmt.Println("Ошибка")
		return
	}
	file, err := myAccount.ToBytes()

	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}

	files.WriteFile(file, "data.json")
}

func getMenu() int {
	var input int
	fmt.Println("Выберите вариант:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scanln(&input)

	return input
}

func findAccount() {
	fmt.Println("Поиск...")
}

func removeAccount() {
	fmt.Println("Удаление...")
}
