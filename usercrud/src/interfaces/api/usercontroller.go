package controllers

import (
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"usercrud/src/domain"
	"usercrud/src/interfaces/database"
	"usercrud/src/usecase"

	"github.com/labstack/echo"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c echo.Context) error {
	u := domain.User{}
	err := c.Bind(&u)
	if err != nil {
		return err
	}
	err = controller.Interactor.Add(u)
	return c.NoContent(errToHttpCode(err))
}

func (controller *UserController) GetUser(c echo.Context) error {
	res, err := controller.Interactor.GetInfo()
	return c.JSON(errToHttpCode(err), res)
}

func (controller *UserController) Delete(c echo.Context, id string) error {
	err := controller.Interactor.Delete(id)
	return c.NoContent(errToHttpCode(err))
}

func (controller *UserController) Update(c echo.Context, idstr string) error {
	u := domain.User{}
	err := c.Bind(&u)
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return err
	}
	u.ID = uint(id)
	err = controller.Interactor.Update(u)
	return c.NoContent(errToHttpCode(err))
}

func errToHttpCode(err error) int {
	switch err {
	case nil:
		return http.StatusOK
	case gorm.ErrInvalidData:
	case gorm.ErrInvalidValue:
		return http.StatusBadRequest
	case gorm.ErrRecordNotFound:
		return http.StatusNotFound
	case gorm.ErrDuplicatedKey:
		return http.StatusConflict
	}
	return http.StatusInternalServerError
}
