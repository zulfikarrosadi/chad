package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zulfikarrosadi/chad/internal/user/helpers"
)

type UserAPI interface {
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type UserAPIImpl struct {
	UserService
}

func NewUserAPI(userService *UserServiceImpl) UserAPI {
	return UserAPIImpl{
		UserService: userService,
	}
}

func (UserAPI UserAPIImpl) Get(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("userID"))
	r := UserAPI.UserService.Get(c.Request().Context(), userID)
	return c.JSON(r.Code, r)
}

func (UserAPI UserAPIImpl) Create(c echo.Context) error {
	file, err := c.FormFile("profilePicture")
	if err != nil {
		return err
	}
	fileName, err := helpers.Upload(file)
	if err != nil {
		return err
	}

	user := &UserCreateRequest{
		Email:          c.FormValue("email"),
		Username:       c.FormValue("username"),
		Password:       c.FormValue("password"),
		ProfilePicture: fileName,
	}

	r := UserAPI.UserService.Create(c.Request().Context(), user)
	return c.JSON(r.Code, r)
}

func (UserAPI UserAPIImpl) Update(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("userID"))

	user := &UserUpdateRequest{
		ID:             i,
		Email:          c.FormValue("email"),
		Username:       c.FormValue("username"),
		Password:       c.FormValue("password"),
		ProfilePicture: c.FormValue("profilePicture"),
	}

	r := UserAPI.UserService.Update(c.Request().Context(), user)
	return c.JSON(r.Code, r)
}

func (UserAPI UserAPIImpl) Delete(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("userID"))
	r := UserAPI.UserService.Delete(c.Request().Context(), userID)
	if r.Code == http.StatusNoContent {
		return c.NoContent(r.Code)
	}

	return c.JSON(r.Code, r)
}
