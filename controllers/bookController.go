package controllers

import (
	"errors"
	"golang-webservice/database"
	"golang-webservice/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooks(ctx *gin.Context) {
	db := database.GetDB()

	var Books []models.Book
	err := db.Find(&Books).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Book not found")
			return
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, Books)
}

func GetBook(ctx *gin.Context) {
	BookID := ctx.Param("bookID")

	db := database.GetDB()

	Book := models.Book{}
	err := db.First(&Book, "id = ?", BookID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Book not found")
			return
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, Book)
}

func CreateBook(ctx *gin.Context) {
	db := database.GetDB()

	type BookInput struct {
		NameBook string `json:"name_book" binding:"required"`
		Author   string `json:"author" binding:"required"`
		UserID   uint   `json:"user_id" binding:"required"`
	}
	var input BookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	Book := models.Book{
		NameBook: input.NameBook,
		Author:   input.Author,
		UserID:   input.UserID,
	}

	err := db.Create(&Book).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, Book)
}

func DeleteBook(ctx *gin.Context) {
	BookID := ctx.Param("bookID")

	db := database.GetDB()

	Book := models.Book{}
	err := db.Where("id = ?", BookID).Delete(&Book).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Book not found")
			return
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}

func UpdateBook(ctx *gin.Context) {
	BookID := ctx.Param("bookID")

	db := database.GetDB()

	type BookInput struct {
		NameBook string `json:"name_book" binding:"required"`
		Author   string `json:"author" binding:"required"`
		UserID   uint   `json:"user_id" binding:"required"`
	}
	var input BookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	Book := models.Book{}
	BookUpdate := models.Book{
		NameBook: input.NameBook,
		Author:   input.Author,
		UserID:   input.UserID,
	}

	err := db.Model(&Book).Where("id = ?", BookID).Updates(BookUpdate).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Book not found")
			return
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, BookUpdate)
}
