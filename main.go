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
	{Id: "2", Name: "Meow"},
}

func RetrieveBooks(c *gin.Context) {
	// for i := range Books {
	// 	fmt.Println(i)
	// }
	c.IndentedJSON(200, Books)
}

func AddBook(c *gin.Context) {
	var NewBook Book
	if err := c.BindJSON(&NewBook); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Wrong format of data"})
		return
	}
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

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var index int = -1

	for i, t := range Books {
		if t.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no book by given Id"})
		return
	}

	var UpdatedBook Book
	if err := c.BindJSON(&UpdatedBook); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Wrong format of Book info"})
		return
	}

	Books[index] = UpdatedBook

	c.IndentedJSON(http.StatusOK, Books)

}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var index int = -1

	for i, t := range Books {
		if t.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Wrong Id of book to delete"})
		return
	}

	Books = append(Books[:index], Books[index+1:]...)

	c.IndentedJSON(http.StatusOK, Books)

}

func main() {
	r := gin.Default()
	r.GET("/Books", RetrieveBooks)
	r.POST("/Books", AddBook)
	r.GET("/Books/:id", GetBook)
	r.PATCH("/Books/:id", UpdateBook)
	r.DELETE("/Books/:id", DeleteBook)
	r.Run("localhost:8080")

}
