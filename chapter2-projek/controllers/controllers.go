package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"chapter2-projek/models"

)

func GetAll(ctx *gin.Context)  {
	Books := []models.Book{}
	db := ctx.MustGet("db").(*gorm.DB)

	tx := db.Find(&Books)
	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, Books)
}

func GetByID(ctx *gin.Context)  {
	id := ctx.Param("id")
	db := ctx.MustGet("db").(*gorm.DB)
	var book models.Book

	err := db.Where("id = ?", id).First(&book).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, book)
}


func CreateBook(ctx *gin.Context)  {
	var newBook models.BookCreated
	db := ctx.MustGet("db").(*gorm.DB)

	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	book := models.Book{
		Title: newBook.Title,
		Author: newBook.Author,
	}

	err = db.Create(&book).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, newBook)

}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	db := ctx.MustGet("db").(*gorm.DB)
	var updatedInput models.BookUpdated
	var book models.Book

	err := ctx.ShouldBindJSON(&updatedInput)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}

	updatedBook := models.Book{
		Title: updatedInput.Title,
		Author: updatedInput.Author,
		UpdatedAt: time.Now(),
	}

	err = db.Model(&book).Where("id = ?", id).Updates(&updatedBook).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}	
	ctx.JSON(http.StatusOK, book)
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	db := ctx.MustGet("db").(*gorm.DB)
	var book models.Book

	err := db.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}	

	ctx.JSON(http.StatusOK, book)
}