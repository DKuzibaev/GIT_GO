package main

import (
	"demo/password/account"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault()

	fmt.Println("___Менеджер паролей___\n")

Menu:
	for {
		variant := getMenu()

		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
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

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин:")
	password := promptData("Введите пароль:")
	url := promptData("Введите URL:")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		fmt.Println("Ошибка")
		return
	}

	vault.AddAccount(*myAccount)
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

func findAccount(vault *account.Vault) {
	url := promptData("Введите ULR адрес для поиска: ")
	accounts := vault.FindAccountByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено!")
	}

	for _, acc := range accounts {
		acc.Output()
	}
}

func removeAccount() {
	fmt.Println("Удаление...")
}
