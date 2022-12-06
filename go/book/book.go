package book

type Book struct {
    Title string
    Author string
    isAvailable bool
}

func NewBook(title string, author string) *Book {
    return &Book{Title: title, Author: author, isAvailable: false}
}

func (b *Book) Lend() bool {
    if b.isAvailable {
        b.isAvailable = false
        return true
    }
    return false
}

func (b *Book) Return() bool {
    b.isAvailable = true
    return true
}


type LibraryBook struct {
    Book *Book
    Department string
}

func NewLibraryBook(title string, author string, department string) *LibraryBook {
    return &LibraryBook{
        Book: &Book{Title: title, Author: author, isAvailable: false},
        Department: department}
}