package main

import (
	"database/sql"
	_	"github.com/mattn/go-sqlite3"
	"fmt"
	"go-api-test/internal/service"
	"go-api-test/internal/store"
	"go-api-test/internal/transport"
	"log"
	"net/http"
)

func main(){

	//connect to sqlite
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//create table if doesnt exists
	q := `CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,	
		author TEXT NOT NULL
	)` 
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}
	//Inject dependencies 
	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHanlder := transport.New(bookService)
	//Config routes
	http.HandleFunc("/books", bookHanlder.HandleBooks)
	http.HandleFunc("/books/", bookHanlder.HandleBookById)

	fmt.Println("🚀Servidor ejecutandose en: http://localhost:8080")
	fmt.Println(" GET     /books      - get all books")
	fmt.Println(" POST    /books      - create new book")
	fmt.Println(" GET     /books/{id} - get books by id")
	fmt.Println(" PUT     /books/{id} - update book")
	fmt.Println(" DELETE  /books/{id} - delete book")
	
	//Listen to server
	log.Fatal(http.ListenAndServe(":8080", nil))
}