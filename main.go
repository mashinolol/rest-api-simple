package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var Books = []Book{
	{Id: 1, Name: "1984"},
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

func main() {
	r := gin.Default()
	r.GET("/Books", RetriveBooks)
	r.POST("/Books", AddBook)
	r.Run("localhost:8080")

}
