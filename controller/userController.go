package controller

import (
	"go-clean-code/contracts/request"
	"go-clean-code/contracts/response"
	"go-clean-code/handlers/command"
	"go-clean-code/handlers/query"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Command command.UserCommand
	Query   query.UserQuery
}

func NewUserController(command command.UserCommand, query query.UserQuery) *UserController {
	return &UserController{
		Command: command,
		Query:   query,
	}
}

func (con *UserController) MapUserEndpoint(router *gin.RouterGroup) {
	router.GET("/users", con.getUsers)
	router.POST("/user", con.createUser)
}

func (con *UserController) getUsers(c *gin.Context) {
	dataUser, err := con.Query.GetUsers()
	if err != nil {
		res := response.ErrorResponse{
			Status:    false,
			ErrorCode: 500,
			Message:   "Internal server error",
		}

		c.JSON(500, res)
		return
	}

	res := response.Response{
		Status:  true,
		Code:    200,
		Message: "Success get data",
		Data:    dataUser,
	}

	c.JSON(200, res)
}

func (con *UserController) createUser(c *gin.Context) {
	var request request.UserCreateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		res := response.ErrorResponse{
			Status:    false,
			ErrorCode: 500,
			Message:   "Failed to bind json",
		}
		c.JSON(500, res)
		return
	}

	err := con.Command.CreateUser(request)
	if err != nil {
		res := response.ErrorResponse{
			Status:    false,
			ErrorCode: 500,
			Message:   "Failed to insert data",
		}
		c.JSON(500, res)
		return
	}

	res := response.Response{
		Status:  true,
		Code:    200,
		Message: "Successfully insert data",
		Data:    nil,
	}

	c.JSON(200, res)
}
