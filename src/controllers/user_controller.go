package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	service "github.com/saefullohmaslul/golang-example/src/services"
	"github.com/saefullohmaslul/golang-example/src/validation"
)

// UserController is controller for user module
type UserController struct {
}

// GetUsers will retrieve all user
func (u UserController) GetUsers(c *gin.Context) {
	users := new(service.UserService).GetUsers()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get all users",
		"result":  users,
	})
}

// GetUser will retrieve user
func (u UserController) GetUser(c *gin.Context) {
	userService := service.UserService{}
	param := validation.GetUserParamSchema{}
	_ = c.ShouldBindUri(&param)

	user := userService.GetUser(int64(param.ID))
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get user",
		"result":  user,
	})
}

// CreateUser will add user into database
func (u UserController) CreateUser(c *gin.Context) {
	userService := service.UserService{}
	var user entity.User
	_ = c.BindJSON(&user)

	data := userService.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success create user",
		"result":  data,
	})
}
