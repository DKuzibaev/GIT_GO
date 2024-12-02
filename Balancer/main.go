package main

import "fmt"

func main() {
	var sum float64
	transactions := transactionsDataBase()
	fmt.Println("Хотите узнать полную сумму транзакций? y/n")
	var input string

	fmt.Scan(&input)

	if input == "y" || input == "Y" {

		for _, value := range transactions {
			sum += value
		}

		fmt.Println(sum)
	} else {
		fmt.Println("Хорошего дня!")
	}

}

func scanTransaction() float64 {
	var userInput float64
	fmt.Println("Введите транзакцию (n выхода): ")
	fmt.Scan(&userInput)
	return userInput
}

func transactionsDataBase() []float64 {
	var transactions = []float64{}

	for {
		value := scanTransaction()
		transactions = append(transactions, value)

		if value == 0 {
			break
		}
	}

	return transactions
}
