package main

import (
"fmt"
"database/sql"
"net/http"
"strconv"

"github.com/gin-gonic/gin"
_ "github.com/lib/pq"

)
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"description"`
}

var mapBooks = make(map[int]Book, 0)
var counter int
var db *sql.DB

func init() {

	var err error
	
	db,err = sql.Open("postgres" , "host = localhost port = 5432 user = postgres password = 301100 dbname = hactiv8-dts-go sslmode = disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database", db)
}

func main() {
	
	g := gin.Default()

	g.GET("/book", getBookHandler)
	g.GET("/book/:id", getBookIdHandler)
	g.POST("/book", addBookHandler)
	g.DELETE("/book/:id", deleteBookHandler)
	g.PUT("/book/:id", updateBookHandler)


	g.Run(":8080")
}

func getBookHandler(ctx *gin.Context) {
	query := "select * from book"

	rows, err := db.Query(query) 
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	books := make([]Book, 0)

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
		})
		return
	}
	books = append(books, book)
}
	ctx.JSON(http.StatusOK, books)
}

func getBookIdHandler(ctx *gin.Context) {
	idString := ctx.Param("id")

	
	id, err := strconv.Atoi(idString)

	var book Book

	query := ("SELECT id, title, author, description FROM book WHERE id = $1")

	row := db.QueryRow(query, id)

	err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)

	if err != nil {
    	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	return
	}
	ctx.JSON(http.StatusOK, book)

}


func addBookHandler(ctx *gin.Context) {
	var newBook Book

	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := "insert into book (title, author, description) values ($1 ,$2 ,$3) returning *"

	row := db.QueryRow(query,newBook.Title, newBook.Author, newBook.Desc)

	err = row.Scan(&newBook.ID, &newBook.Title, &newBook.Author, &newBook.Desc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, newBook)
}

func deleteBookHandler(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var deletedBook Book

	query := "delete from book where id =$1 "

	row := db.QueryRow(query, id)
	err = row.Scan(&deletedBook.ID, &deletedBook.Title, &deletedBook.Author, &deletedBook.Desc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, deletedBook)	

}

func updateBookHandler(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updatedBook Book

	ctx.ShouldBindJSON(&updatedBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	query := "update book set title=$1, author=$2, description=$3 where id=$4 returning id"

	row := db.QueryRow(query, updatedBook.Title, updatedBook.Author, updatedBook.Desc, id)

	err = row.Scan(&updatedBook.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedBook)
}
