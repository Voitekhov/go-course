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
	// &account - адрес переменной account в памяти 0xc000092108
	fmt.Println("---------- Вопрос №2 ----------")
	// Что вывелось в консоль?
	// Почему первый первый и второй вывод отличаются?
	account2 := account
	fmt.Printf("%p \n", &account2)
	fmt.Printf("%p \n", &account)
	// При создании новой переменной, хоть и равной первой для нее будет выделено новое место в памяти
	fmt.Println("---------- Вопрос №3 ----------")
	// Какие значения будут храниться в переменных a b c?
	// Объясните каждый из выводов в консоль
	a := 2  // 3
	b := &a // 0xc00008c060
	c := &a // 0xc00008c060
	a++
	fmt.Println(&a, &b, &c) // &a адрес переменной, b и c хранят адрес переменной a. При этом мы запрашиваем адреса b и c, это отдельные переменные и имеют свои адреса
	fmt.Println(a, b, c)    // a=3, т.к. при выполнении кода она получила значение a++=3, b и с все еще хранят в себе адрес переменной a
	fmt.Println(a, *b, *c)  // *b  и *с * выводит значение которое хранится по указанному адресу, соответствуют - 3
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
}

// a - переменная цикла, при каждой иттерацию в нее копируется значение, каждый раз мы добавляем указатель на временную переменную
// при каждом вызове a - новый временный объект в куче, поэтому адреса меняются между вызовами
