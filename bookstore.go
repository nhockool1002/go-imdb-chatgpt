package main

import (
    "fmt"
)

type Book struct {
    Title     string
    Author    string
    ISBN      string
    Quantity  int
}

type Bookstore struct {
    books []Book
}

func (store *Bookstore) addBook(book Book) {
    store.books = append(store.books, book)
}

func (store Bookstore) showBooks() {
    fmt.Println("Danh sách sách trong cửa hàng:")
    for _, book := range store.books {
        fmt.Printf("Tiêu đề: %s, Tác giả: %s, ISBN: %s, Số lượng: %d\n", book.Title, book.Author, book.ISBN, book.Quantity)
    }
}

func (store *Bookstore) searchByTitle(title string) []Book {
    var results []Book
    for _, book := range store.books {
        if book.Title == title {
            results = append(results, book)
        }
    }
    return results
}

func (store *Bookstore) sellBook(title string) {
    for i, book := range store.books {
        if book.Title == title {
            if store.books[i].Quantity > 0 {
                store.books[i].Quantity--
                fmt.Printf("Bạn đã mua một cuốn sách có tiêu đề \"%s\"\n", title)
                return
            } else {
                fmt.Printf("Xin lỗi, sách có tiêu đề \"%s\" đã hết hàng\n", title)
                return
            }
        }
    }
    fmt.Printf("Xin lỗi, không tìm thấy sách có tiêu đề \"%s\"\n", title)
}

func main() {
    var myBookstore Bookstore

    // Thêm một số sách vào cửa hàng
    myBookstore.addBook(Book{Title: "To Kill a Mockingbird", Author: "Harper Lee", ISBN: "9780061120084", Quantity: 5})
    myBookstore.addBook(Book{Title: "Pride and Prejudice", Author: "Jane Austen", ISBN: "9780141439518", Quantity: 3})
    myBookstore.addBook(Book{Title: "1984", Author: "George Orwell", ISBN: "9780451524935", Quantity: 7})

    // Hiển thị danh sách sách trong cửa hàng
    myBookstore.showBooks()

    // Tìm kiếm sách theo tiêu đề
    fmt.Println("\nTìm kiếm sách có tiêu đề \"1984\":")
    searchResults := myBookstore.searchByTitle("1984")
    for _, book := range searchResults {
        fmt.Printf("Tiêu đề: %s, Tác giả: %s, ISBN: %s, Số lượng: %d\n", book.Title, book.Author, book.ISBN, book.Quantity)
    }

    // Bán một cuốn sách
    myBookstore.sellBook("To Kill a Mockingbird")
    myBookstore.showBooks()
}
