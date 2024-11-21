package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const IMTpower = 2

func main() {

	var helloMsg = "__Калькулятор индекса массы тела__\n"
	fmt.Println(helloMsg)
	var userHeight float64
	var userWeigth float64
	i := true
	for i {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Хотите произвести новый расчёт?(Y/N)")
		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.TrimSpace(userAnswer)

		if userAnswer == "n" || userAnswer == "N" {
			fmt.Println("Хорошего дня!")
			i = false
			break
		} else {
			userHeight, userWeigth = getUserInput()
			IMT := calculateIMT(userWeigth, userHeight)
			fmt.Println(outputResult(IMT))
			checkIMTNorm(IMT)
		}
	}
}

func waitForExit() {
	fmt.Println("\nХорошего Дня...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func outputResult(IMT float64) string {
	result := fmt.Sprintf("Индекс массы тела равен: %.2f\n", IMT)
	return result
}

func calculateIMT(userWeigth, userHeight float64) float64 {
	result := userWeigth / math.Pow(userHeight/100, IMTpower)
	return result
}

func checkIMTNorm(IMT float64) {
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
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeigth float64

	fmt.Println("Введите свой рост в сантиметрах:")
	_, err := fmt.Scanln(&userHeight)
	if err != nil || userHeight <= 0 {
		fmt.Println("Ошибка: Некорректный ввод роста.")
		waitForExit()
		return 0.0, 0.0
	}

	fmt.Println("Введите свой вес:")
	_, err = fmt.Scanln(&userWeigth)
	if err != nil || userWeigth <= 0 {
		fmt.Println("Ошибка: Некорректный ввод веса.")
		waitForExit()
		return 0.0, 0.0
	}

	return userHeight, userWeigth
}
