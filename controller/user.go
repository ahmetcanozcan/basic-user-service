package controller

import (
	"app/common/herrors"
	"app/model"
	"app/service"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (ct *UserController) CreateUser(c echo.Context) error {

	ctx := ct.bindContext(c)

	request := new(CreateUserRequest)

	if err := c.Bind(request); err != nil {
		return herrors.NewBadRequest(err)
	}

	requestedUser := &model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := ct.userService.Create(ctx, requestedUser)

	if err != nil {
		return err
	}

	response := &CreateUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(http.StatusOK, response)
}

func (ct *UserController) EditUser(c echo.Context) error {
	ctx := ct.bindContext(c)

	request := new(EditUserRequest)

	if err := c.Bind(request); err != nil {
		return herrors.NewBadRequest(err)
	}

	if err := ct.userService.EditUser(ctx, request.ID, request.Name, request.Password); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (ct *UserController) GetUsers(c echo.Context) error {
	ctx := ct.bindContext(c)

	users, err := ct.userService.GetUsers(ctx)

	if err != nil {
		return err
	}

	response := mapUserListResponse(users)

	return c.JSON(http.StatusOK, response)
}

func (ct *UserController) GetUser(c echo.Context) error {
	ctx := ct.bindContext(c)

	request := new(GetUserRequest)

	if err := c.Bind(request); err != nil {
		return herrors.NewBadRequest(err)
	}

	user, err := ct.userService.GetUserByID(ctx, request.ID)

	if err != nil {
		return err
	}

	response := mapUserResponse(user)

	return c.JSON(http.StatusOK, response)
}

func (ct *UserController) DeleteUser(c echo.Context) error {
	ctx := ct.bindContext(c)

	request := new(DeleteUserRequest)

	if err := c.Bind(request); err != nil {
		return herrors.NewBadRequest(err)
	}

	if err := ct.userService.DeleteUser(ctx, request.ID); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (ct *UserController) bindContext(c echo.Context) context.Context {
	return c.Request().Context()
}

func mapUserResponse(user *model.User) UserListResponseItem {
	return UserListResponseItem{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func mapUserListResponse(users []*model.User) []UserListResponseItem {
	var response []UserListResponseItem
	for _, user := range users {
		response = append(response, mapUserResponse(user))
	}
	return response
}

func MakeUserHandler(e *echo.Echo, ct *UserController) {
	e.PUT("/users", ct.CreateUser)
	e.PATCH("/users/:id", ct.EditUser)
	e.GET("/users", ct.GetUsers)
	e.GET("/users/:id", ct.GetUser)
	e.DELETE("/users/:id", ct.DeleteUser)
}
