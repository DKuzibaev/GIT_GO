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
	fmt.Println("__Калькулятор индекса массы тела__\n")

	for {
		userHeight, userWeigth := getUserInput()
		IMT := calculateIMT(userWeigth, userHeight)
		fmt.Println(outputResult(IMT))
		checkIMTNorm(IMT)

		isRepeateCalc := checkRepeatCalc()

		if !isRepeateCalc {
			break
		}
	}

}

func waitForExit() {
	fmt.Println("\nХорошего Дня...")
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
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeigth float64

	fmt.Println("Введите свой рост в сантиметрах:")
	_, err := fmt.Scanln(&userHeight)
	if err != nil || userHeight <= 0 {
		panic("Ошибка: Некорректный ввод роста.")
	}

	fmt.Println("Введите свой вес:")
	_, err = fmt.Scanln(&userWeigth)
	if err != nil || userWeigth <= 0 {
		panic("Ошибка: Некорректный ввод веса.")
	}

	return userHeight, userWeigth
}

func checkRepeatCalc() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Хотите произвести новый расчёт?(Y/N)")
	userAnswer, _ := reader.ReadString('\n')
	userAnswer = strings.TrimSpace(userAnswer)

	if userAnswer == "y" || userAnswer == "Y" {
		return true
	}
	waitForExit()
	return false
}
