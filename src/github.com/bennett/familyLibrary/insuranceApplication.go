package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/bennett/familyLibrary/endpoint"
)

func main() {
	//println("Hello Gin")
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.Static("/statics", "statics")

	// index page
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{"body": "Insurance Book"})
	})

	// Family Library
	// book page
	/*router.GET("/book/list", func(context *gin.Context) {
		context.HTML(http.StatusOK, "book-list.html", gin.H{"title": "Book List"})
	})*/
	book := router.Group("/book")
	{
		book.GET("/list", endpoint.BookListEndpoint)
		book.GET("/detail", endpoint.BookDetailEndpoint)
	}

	router.Run(":8080")
}

