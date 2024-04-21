package infrastructure

import (
	controllers "usercrud/src/interfaces/api"
	"usercrud/src/interfaces/database"

	"github.com/labstack/echo"
)

func Init(sqlHandler database.SqlHandler) {
	// Echo instance
	e := echo.New()
	userController := controllers.NewUserController(sqlHandler)

	e.GET("/users", func(c echo.Context) error {
		return userController.GetUser(c)
	})

	e.POST("/users", func(c echo.Context) error {
		return userController.Create(c)
	})

	e.DELETE("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		return userController.Delete(c, id)
	})

	e.PUT("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		return userController.Update(c, id)
	})

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
