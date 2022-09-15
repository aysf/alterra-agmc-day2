package routes

import (
	"day2/controllers/books"
	"day2/controllers/users"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/v1")

	v1Users := v1.Group("/users")

	v1Users.GET("", users.List)
	v1Users.POST("", users.Create)
	v1Users.GET("/:id", users.Show)
	v1Users.PUT("/:id", users.Update)
	v1Users.DELETE("/:id", users.Delete)

	v1Books := v1.Group("/books")

	v1Books.GET("", books.List)
	v1Books.POST("", books.Create)
	v1Books.GET("/:id", books.Show)
	v1Books.PUT("/:id", books.Update)
	v1Books.DELETE("/:id", books.Delete)

	return e
}
