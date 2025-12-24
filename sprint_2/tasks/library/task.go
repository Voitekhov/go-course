package library

import (
	"fmt"
	"sort"
	"time"
)

/*
Задания:
 1. Реализовать функцию addBook, которая добавляет в библиотеку книгу. Дубликатов по id быть не должно.
 2. Написать функцию, которая показывает наличие книг/книги по названию в библиотеке
 3. Расширить type Book struct и добавить информацию о пользователях, которые читали книгу
 4. Написать функцию, которая помечает что пользователь Вася Пупкин прочитал книгу с ID в такую-то дату
 5. Написать функцию, которая возвращает все книги за заданый год
 6. Написать функцию, которая возвращает все книги определенного Автора
 7. Написать функцию, которая сортирует книги по году выпуска
*/
type Book struct {
	ID      int
	Title   string
	Author  string
	Year    int
	Pages   int
	Readers []ReadingInfo // 3
}
type ReadingInfo struct {
	User string
	Date time.Time
}
type Library struct {
	Name  string
	Books []Book
}

// 1
func (l *Library) AddBook(book Book) error {
	for _, b := range l.Books {
		if b.ID == book.ID {
			return fmt.Errorf("book with ID %d already exists", book.ID)
		}
	}
	l.Books = append(l.Books, book)
	return nil
}

//2

func (l *Library) FindBookByTitle(Title string) []Book {
	var result []Book
	for _, b := range l.Books {
		if b.Title == Title {
			result = append(result, b)
		}
	}
	return result
}

// 4
func (l *Library) MarkAsRead(bookID int, user string, date time.Time) error {
	for i := range l.Books {
		if l.Books[i].ID == bookID {

			readers := l.Books[i].Readers

			for _, r := range readers {
				if r.User == user {
					return fmt.Errorf("user already marked as read")
				}
			}

			l.Books[i].Readers = append(
				readers,
				ReadingInfo{
					User: user,
					Date: date,
				},
			)
			return nil
		}
	}
	return fmt.Errorf("book not found")
}

// 5
func (l *Library) getBooksByYear(Year int) []Book {
	var result []Book
	for _, b := range l.Books {
		if b.Year == Year {
			result = append(result, b)
		}
	}
	return result
}

// 6
func (l *Library) FindByAuthor(Author string) []Book {
	var result []Book
	for _, b := range l.Books {
		if b.Author == Author {
			result = append(result, b)
		}
	}
	return result
}

// 7
func (l *Library) SortByYear() {
	sort.Slice(l.Books, func(i, j int) bool {
		return l.Books[i].Year < l.Books[j].Year
	})

}
