package handlers

import (
	"exemplo/GO-API/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListBooks(context echo.Context) error {
	var books []models.Book
	if err := models.DB.Find(&books).Error; err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, books)
	return nil

}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(context echo.Context) error {
	// Valida a entrada
	var input CreateBookInput
	if err := context.Bind(&input); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	book := models.Book{Title: input.Title, Author: input.Author}

	err := models.DB.Create(&book).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusCreated, echo.Map{
		"Livro criado com sucesso": book,
	})
	return nil
}

//Localizar ID do livro
//Achar um livro

func FindBook(context echo.Context) error {
	// Achar um modelo caso exista
	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	context.JSON(http.StatusOK, book)
	return nil
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

//PATCH /books/:id
// Atualiza um livro

func UpdateBook(context echo.Context) error {
	// Acha um modelo caso exista

	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())

	}

	// Validar input
	var input UpdateBookInput

	if err := context.Bind(&input); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{
			"Erro encontrado": err.Error(),
		})

	}

	UpdateBook := models.Book{Title: input.Title, Author: input.Author}

	if err := models.DB.Model(&book).Where("id = ?", context.Param("id")).Updates(&UpdateBook).Error; err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, book)
	return nil
}

// Deleta um registro de livro

func DeleteBook(context echo.Context) error {
	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	if err := models.DB.Delete(&book); err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}
	context.JSON(http.StatusOK, "Registro deletado")
	return nil
}
