package main

import "fmt"

func main() {
	transactions := transactionsDataBase()
	fmt.Println(calculateTransactions(transactions))

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

func calculateTransactions(transactions []float64) float64 {
	fmt.Println("Хотите узнать полную сумму транзакций? y/n")

	var input string
	var sum float64

	fmt.Scan(&input)

	if input == "y" || input == "Y" {

		for _, value := range transactions {
			sum += value
		}
		fmt.Printf("Ваш баланс: %.2f", sum)
	} else {
		fmt.Println("Хорошего дня!")
	}

	return sum
}
