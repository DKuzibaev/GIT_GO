package main

import "fmt"

// Добавил alias для получаемого и возвращаемого типа!
type bookmarksMap = map[string]string

//Глобальная переменная для хранения данных
var database = bookmarksMap{}

func main() {
	fmt.Println("Приложение для закладок URL адресов\n")
	// Добавил lable Menu: для остановки цикла!
Menu:
	for {
		variant := getMenu()

		switch variant {
		case 1:
			showDBItems(database)
		case 2:
			addToDB(database)
		case 3:
			removeFromDB(database)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var input int
	fmt.Println("Выберите вариант:")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход\n")
	fmt.Scan(&input)

	return input
}

func addToDB(db bookmarksMap) bookmarksMap {
	var key, val string
	fmt.Println("Введите название: ")
	fmt.Scan(&key)
	fmt.Println("Введите адрес: ")
	fmt.Scan(&val)
	db[key] = val

	return db
}

func showDBItems(db bookmarksMap) {
	if len(db) == 0 {
		fmt.Println("Тут пусто...")
	}
	for key, val := range db {
		fmt.Printf("%s: %s\n", key, val)
	}
}

func removeFromDB(db bookmarksMap) {
	var id string
	fmt.Println("Введите название")
	fmt.Scan(&id)

	fmt.Printf("%s успешно удален..\n", id)
	delete(db, id)
}
