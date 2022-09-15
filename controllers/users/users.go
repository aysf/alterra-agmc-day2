package users

import (
	"day2/lib/database"
	"day2/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success get all users",
		"users":   users,
	})
}

func Show(c echo.Context) error {

	id := c.Param("id")
	user, err := database.FindUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success get user",
		"user":    user,
	})
}

func Create(c echo.Context) error {

	var user models.User

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := database.NewUser(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success create new user",
	})
}

func Update(c echo.Context) error {

	id := c.Param("id")

	var user models.User

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user.ID = uint(userId)

	if _, err := database.UpdateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success update user",
	})
}

func Delete(c echo.Context) error {

	id := c.Param("id")

	if err := database.DeleteUser(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success delete user",
	})
}
