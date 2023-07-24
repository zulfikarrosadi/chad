package main

import (
	"github.com/labstack/echo/v4"
	"github.com/zulfikarrosadi/chad/internal/user"
	"github.com/zulfikarrosadi/chad/pkg/db"
)

func main() {
	e := echo.New()
	userRepository := user.NewUserRepository(db.GetConnection())
	userService := user.NewUserService(userRepository.(user.UserRepositoryImpl))
	userHandler := user.NewUserAPI(userService.(*user.UserServiceImpl))

	e.POST("/api/users", userHandler.Create)
	e.GET("/api/users/:userID", userHandler.Get)
	e.PUT("/api/users/:userID", userHandler.Update)
	e.DELETE("/api/users/:userID", userHandler.Delete)

	e.Logger.Fatal(e.Start("localhost:3000"))
}
