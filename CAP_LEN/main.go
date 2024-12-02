package main

import "fmt"

func main() {
	//tr := [5]int{1, 2, 3, 4, 5}
	//В таком случае исходный массив не меняется!!
	// trNew := tr
	// trNew[0] = 30

	//fmt.Println(trNew)

	//В таком случае исходный массив меняется так как тут передается ссылка на основной массив!!
	//Слайс - это по сути окно которое ссылается на основные элементы, получается любое изменение в срезе
	//меняется и исходгый элемент
	// trPart := tr[1:]
	// trPart[0] = 10
	// fmt.Println(trPart)
	// fmt.Println(tr)

	// trPartNew := trPart[:1]
	// fmt.Println(trPartNew)

	// //свойство slice
	// // len - длинная среза, cap - вместимость среза
	// fmt.Println(len(trPart), cap(trPart))
	// fmt.Println(len(trPartNew), cap(trPartNew))

	//Динамический массив
	// var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// slice := arr[:2]
	// slice[0] = -1
	// fmt.Println(arr)
	// fmt.Println(slice)

	// //Добавление элементов
	// newArr := append(arr, 0)
	// newArr[0] = 1
	// fmt.Println(newArr)

	// 1. Кол-во транзакций
	// 2. Добавлять в массив транзакцийй
	// 3. Вывести массив

	fmt.Println("Введите количество транзакциий")
	var count int
	trans := []int{}
	fmt.Scan(&count)

	for i := 1; i <= count; i++ {
		var userInput int
		var stop string
		fmt.Printf("Введите сумму %d\n", i)
		fmt.Scan(&userInput)
		trans = append(trans, userInput)
		fmt.Printf("Тразакция №%d, внесено %d\n", i, trans[i-1])

		if count == i {
			fmt.Println("Хотите продолжить? y/n")
			fmt.Scan(&stop)

			if stop == "y" || stop == "Y" {
				count = count * count
				continue
			} else {
				break
			}
		}
	}

	var result string
	fmt.Println("Хотите узнать итоговую сумму?")
	fmt.Scan(&result)

	if result == "y" || result == "Y" {
		totalSum := 0
		for i := 0; i < len(trans); i++ {
			totalSum += trans[i]
		}
		fmt.Printf("Итоговая сумма равно: %d", totalSum)
	} else {
		fmt.Println("Всего хорошего..")
	}

}
