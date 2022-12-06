package main

import (
    "fmt"
    book "examples/book"
)

func main() {
    b := book.NewBook("Harry Potter and the Philosopher's Stone",
                      "J.K. Rowling")

    fmt.Printf("New book created. Title: %s Author: %s\n",
               b.Title, b.Author)

    if b.Lend() {
        fmt.Println("Book successfully lent.")
    } else {
        fmt.Println("Book is not currently available. Try in some time.")
    }

    // This doesn't work, unexported field
//     if b.isAvailable {
//         fmt.Println("Book available")
//     }
    lb := book.NewLibraryBook("Harry Potter and the Philosopher's Stone",
                      "J.K. Rowling",
                      "Children's Library, Oak Street")
    fmt.Printf("New book created. Title: %s Author: %s, Department: %s\n",
               lb.Book.Title, lb.Book.Author, lb.Department)
}
