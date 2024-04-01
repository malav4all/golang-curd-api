package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Authour  string `json:"authour"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The Lost", Authour: "Malav", Quantity: 1},
	{ID: "2", Title: "The Lost 1", Authour: "Atul", Quantity: 12},
	{ID: "3", Title: "The Lost 3", Authour: "Reactjs", Quantity: 14},
	{ID: "4", Title: "The Lost 4", Authour: "Javascript", Quantity: 16},
	{ID: "5", Title: "The Lost 6", Authour: "Malav Naagar", Quantity: 19},
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookId(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookId(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("books Not Found")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createdBooks(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	//Created a Router
	router := gin.Default()
	// Make a Get Request
	router.GET("/books", getBooks)

	//Make Route Get BookBy ID

	router.GET("/books/:id", bookById)

	//Make a Post Request
	router.POST("/books", createdBooks)
	// Make  a Listen and Serve request
	router.Run("localhost:8080")
}
