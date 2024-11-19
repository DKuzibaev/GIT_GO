package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	fmt.Println("__Калькулятор индекса массы тела__\n")

	const IMTpower = 2
	var userHeight float64
	var userWeigth float64

	fmt.Println("Введите свой рост в метрах:")
	_, err := fmt.Scanln(&userHeight)
	if err != nil || userHeight <= 0 {
		fmt.Println("Ошибка: Некорректный ввод роста.")
		waitForExit()
		return
	}

	fmt.Println("Введите свой вес:")
	_, err = fmt.Scanln(&userWeigth)
	if err != nil || userWeigth <= 0 {
		fmt.Println("Ошибка: Некорректный ввод веса.")
		waitForExit()
		return
	}

	IMT := userWeigth / math.Pow(userHeight, IMTpower)
	fmt.Printf("Индекс массы тела равен: %.2f\n", IMT)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Произвести расчёт индекса массы тела, который определяет в каком соотношении находятся вес и рост? (Y/N)")

	userAnswer, _ := reader.ReadString('\n')
	userAnswer = strings.TrimSpace(userAnswer)

	if strings.ToUpper(userAnswer) == "Y" {
		if IMT < 16 {
			fmt.Println("Выраженный дефицит массы тела")
		} else if IMT <= 18.5 {
			fmt.Println("Недостаточная (дефицит) масса тела")
		} else if IMT <= 25 {
			fmt.Println("Норма")
		} else if IMT <= 30 {
			fmt.Println("Избыточная масса тела (предожирение)")
		} else if IMT <= 35 {
			fmt.Println("Ожирение первой степени")
		} else {
			fmt.Println("Ожирение третьей степени (морбидное)")
		}
	} else {
		fmt.Println("Хорошего дня!")
	}

	waitForExit()
}

func waitForExit() {
	fmt.Println("\nНажмите Enter для выхода...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
