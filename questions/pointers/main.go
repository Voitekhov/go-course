package main

import "fmt"

type Account struct {
	Id   int64
	Name string
}

var accounts = []Account{
	{
		Id:   1,
		Name: "first",
	},
	{
		Id:   2,
		Name: "second",
	},
	{
		Id:   3,
		Name: "third",
	},
	{
		Id:   4,
		Name: "fourth",
	},
}

func GetAccountsPointers() []*Account {
	res := make([]*Account, 0, len(accounts))
	for _, a := range accounts {
		res = append(res, &a)
	}
	return res
}

func main() {
	fmt.Println("---------- Вопрос №1 ----------")
	// Что вывелось в консоль?
	account := Account{Id: 23, Name: "Mike"}
	fmt.Printf("%p \n", &account)

	fmt.Println("---------- Вопрос №2 ----------")
	// Что вывелось в консоль?
	// Почему первый первый и второй вывод отличаются?
	account2 := account
	fmt.Printf("%p \n", &account2)
	fmt.Printf("%p \n", &account)

	fmt.Println("---------- Вопрос №3 ----------")
	// Почему в консоль постоянно выводятся различне адреса памяти,
	// хотя при каждом вызове функции GetAccountsPointers мы работаем с одним и тем же accounts?

	// Вызов a1
	pointers := GetAccountsPointers()
	for _, pointer := range pointers {
		fmt.Printf("%p \n", pointer)
	}
	fmt.Println()
	// Вызов a2
	pointers = GetAccountsPointers()
	for _, pointer := range pointers {
		fmt.Printf("%p \n", pointer)
	}

}
