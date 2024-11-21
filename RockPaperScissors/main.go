package main

import (
	"fmt"
	"math/rand"
)

func main() {
	start()
}

func start() {
	fmt.Println("Добро пожаловать в игру «Камень, Ножницы, Бумага»")
	userChoice := choosingSubject()
	computerSelection := chooseComputerSelection()

	fmt.Printf("Вы выбрали %s\n", userChoice)
	fmt.Printf("Компьютер выбрал %s\n", computerSelection)

	fmt.Println(game(userChoice, computerSelection))
}

func choosingSubject() string {
	println("Выберите предмет (введите цифру): ")

	arr := [3]string{"Камень", "Ножницы", "Бумага"}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d. %s\n", i+1, arr[i])
	}
	return howIs()
}

func howIs() string {
	var userInput string
	fmt.Scan(&userInput)

	if userInput == "1" {
		return "Камень"
	} else if userInput == "2" {
		return "Ножницы"
	} else if userInput == "3" {
		return "Бумага"
	} else {
		return "Ошибка"
	}
}

func chooseComputerSelection() string {
	arr := [3]string{"Камень", "Ножницы", "Бумага"}
	selection := rand.Int() % 3
	return arr[selection]
}

func game(usV string, cpV string) string {
	arr := [3]string{"Камень", "Ножницы", "Бумага"}

	if usV == cpV {
		return "Ничья"
	} else if usV == arr[0] && cpV == arr[1] {
		return ("Камень бьет Ножницы")
	} else if (usV == arr[1] && cpV == arr[2]) || (usV == arr[2] && cpV == arr[1]) {
		return ("Ножницы бьют Бумагу")
	} else if usV == arr[2] && cpV == arr[0] {
		return ("Бумага бьёт Камень")
	}
	return "ooPs"
}
