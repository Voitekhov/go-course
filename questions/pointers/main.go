package main

import (
	"fmt"
)

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
	// Какие значения будут храниться в переменных a b c?
	// Объясните каждый из выводов в консоль
	a := 2
	b := &a
	c := &a
	a++
	fmt.Println(&a, &b, &c)
	fmt.Println(a, b, c)
	fmt.Println(a, *b, *c)

	fmt.Println("---------- Вопрос №4 ----------")
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

	fmt.Println("---------- Вопрос №5 ----------")
	// Почему в sliceOfSlice 9 вывелась, а в copyOfS нет?
	// почему используется s[:], а не s?
	// если у sliceOfSlice и copyOfS будет тип [][]int, как сделать чтобы 9 была только в sliceOfSlice?
	sliceOfSlice := [][2]int{{1, 2}, {3, 4}, {5, 6}}
	var copyOfS [][]int
	for _, s := range sliceOfSlice {
		copyOfS = append(copyOfS, s[:])
	}
	sliceOfSlice[0][0] = 9
	fmt.Println(sliceOfSlice)
	fmt.Println(copyOfS)
}
