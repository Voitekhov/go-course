package library

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func testLibrary() *Library {
	return &Library{
		Name: "Test Library",
		Books: []Book{
			{
				ID:     1,
				Title:  "Grokking Algorithms",
				Author: "Aditya Bhargava",
				Year:   2016,
				Pages:  278,
				Readers: []ReadingInfo{
					{
						User: "Anton",
						Date: time.Date(2020, time.March, 10, 0, 0, 0, 0, time.Local),
					},
					{
						User: "Sergey",
						Date: time.Date(2025, time.December, 10, 0, 0, 0, 0, time.Local),
					},
				},
			},
			{
				ID:     2,
				Title:  "Design Patterns",
				Author: "Head First",
				Year:   2022,
				Pages:  633,
				Readers: []ReadingInfo{
					{
						User: "Sergey",
						Date: time.Date(2025, time.December, 25, 0, 0, 0, 0, time.Local),
					},
				},
			},
			{
				ID:     3,
				Title:  "1984",
				Author: "George Orwell",
				Year:   1949,
				Pages:  210,
			},
			{
				ID:     4,
				Title:  "The Little Prince",
				Author: "Antoine de Saint-Exupery",
				Year:   1943,
				Pages:  60,
			},
			{
				ID:     5,
				Title:  "Crime and Punishment",
				Author: "Fyodor Dostoevsky",
				Year:   1886,
				Pages:  321,
				Readers: []ReadingInfo{
					{
						User: "Vasya Pupkin",
						Date: time.Date(2007, time.October, 30, 0, 0, 0, 0, time.Local),
					},
				},
			},
		},
	}
}

func TestAddBook(t *testing.T) {
	lib := testLibrary()

	err := lib.AddBook(Book{
		ID:     10,
		Title:  "New Book",
		Author: "Someone",
		Year:   2024,
	})

	require.NoError(t, err)
	require.Len(t, lib.Books, 6)
}

func TestAddBook_DuplicateID(t *testing.T) {
	lib := testLibrary()

	err := lib.AddBook(Book{ID: 1})
	require.Error(t, err)
}

func TestFindBookByTitle(t *testing.T) {
	lib := testLibrary()

	result := lib.FindBookByTitle("1984")

	require.Len(t, result, 1)
	require.Equal(t, "George Orwell", result[0].Author)
}

func TestMarkAsRead(t *testing.T) {
	lib := testLibrary()
	date := time.Now()

	err := lib.MarkAsRead(1, "Ivan", date)

	require.NoError(t, err)
	require.Len(t, lib.Books[0].Readers, 3)
	require.Equal(t, "Ivan", lib.Books[0].Readers[2].User)
}

func TestMarkAsRead_AlreadyRead(t *testing.T) {
	lib := testLibrary()

	err := lib.MarkAsRead(1, "Sergey", time.Now())

	require.Error(t, err)
	require.EqualError(t, err, "user already marked as read")
}

func TestMarkAsRead_NotFound(t *testing.T) {
	lib := testLibrary()

	err := lib.MarkAsRead(999, "Ivan", time.Now())
	require.Error(t, err)
}

func TestBookByYear(t *testing.T) {
	lib := testLibrary()

	result := lib.BookByYear(2016)

	require.Len(t, result, 1)
	require.Equal(t, "Grokking Algorithms", result[0].Title)
}

func TestFindByAuthor(t *testing.T) {
	lib := testLibrary()

	result := lib.FindByAuthor("George Orwell")

	require.Len(t, result, 1)
	require.Equal(t, "1984", result[0].Title)
}

func TestSortByYear(t *testing.T) {
	lib := testLibrary()

	lib.SortByYear()

	require.True(t, lib.Books[0].Year <= lib.Books[1].Year)
	require.True(t, lib.Books[len(lib.Books)-2].Year <= lib.Books[len(lib.Books)-1].Year)
}
