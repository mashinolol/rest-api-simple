package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var Books = []Book{
	{Id: "1", Name: "1984"},
}

func RetriveBooks(c *gin.Context) {
	// for i := range Books {
	// 	fmt.Println(i)
	// }
	c.IndentedJSON(200, Books)
}

func AddBook(c *gin.Context) {
	var NewBook Book
	c.BindJSON(&NewBook)
	Books = append(Books, NewBook)
	c.IndentedJSON(http.StatusCreated, NewBook)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)

}

func GetBookById(id string) (*Book, error) {
	for i, t := range Books {
		if t.Id == id {
			return &Books[i], nil
		}
	}

	return nil, errors.New("to do not found")
}

func main() {
	r := gin.Default()
	r.GET("/Books", RetriveBooks)
	r.POST("/Books", AddBook)
	r.GET("/Books/:id", GetBook)
	r.Run("localhost:8080")

}
