package main


import "github.com/gin-gonic/gin"
import "net/http"
import "strconv"


type Book struct {
	ID int			`json:"id"`
	Title string	`json:"title"`
	Author string	`json:"author"`
	Desc string		`json:"desc"`
}

var mapBooks = make(map[int]Book, 0)
var counter int

func main() {
	g := gin.Default()

	g.GET("/book",getBookHandler)
	g.POST("/book",addBookHandler)
	g.DELETE("/book/:id",deleteBookHandler)
	g.PUT("/book/:id",updateBookHandler)

	g.Run(":8080")
}	

func getBookHandler (ctx *gin.Context)  {
	books := make([]Book, 0)

	for _, v:= range mapBooks{
		books = append(books, v)
	}

	ctx.JSON(http.StatusOK, books)
}

func addBookHandler (ctx *gin.Context)  {
	var newBook Book
	
	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return 
	}

	newBook.ID = counter
	mapBooks [counter] = newBook
	counter++

	ctx.JSON(http.StatusOK, newBook)
}

func deleteBookHandler (ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	v, found := mapBooks[id]
	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}
		delete(mapBooks, id)

		ctx.JSON(http.StatusOK, v)

}

func updateBookHandler (ctx *gin.Context)  {
	idString := ctx.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	_, found := mapBooks[id]
	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}
	var updatedBook Book
	ctx.ShouldBindJSON(&updatedBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}	

	mapBooks[id] = updatedBook

	ctx.JSON(http.StatusOK, updatedBook)
}