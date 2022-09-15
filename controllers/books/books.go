package books

import (
	"day2/config"
	"day2/lib/database"
	"day2/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var DB = config.DB

func List(c echo.Context) error {

	var books []models.Book

	if err := DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success get all users",
		"books":   books,
	})
}

func Show(c echo.Context) error {

	id := c.Param("id")
	book, err := database.FindBookById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success get book",
		"book":    book,
	})
}

func Create(c echo.Context) error {

	var book models.Book

	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := database.NewBook(book); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success create new book",
	})
}

func Update(c echo.Context) error {

	id := c.Param("id")

	var book models.Book

	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	bookId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book.ID = uint(bookId)

	if _, err := database.UpdateBook(book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success update book",
	})
}

func Delete(c echo.Context) error {

	id := c.Param("id")

	if err := database.DeleteBook(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success delete book",
	})
}
