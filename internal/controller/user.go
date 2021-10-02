package controller

import (
	"fmt"
	"github.com/boliev/wordy/internal/domain"
	"github.com/boliev/wordy/internal/request"
	"github.com/boliev/wordy/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// User struct
type User struct {
	userRepository domain.UserRepository
}

// CreateUserController controller for /users/
func CreateUserController(userRepository domain.UserRepository) *User {
	return &User{
		userRepository: userRepository,
	}
}

// List action for [get] /users/
func (u User) List(c *gin.Context) {
	c.JSON(http.StatusOK, response.CreateUsersListFromDomain(u.userRepository.FindAll()))
}

// One action for [get] /users/:id
func (u User) One(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := u.userRepository.Find(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewError(err.Error()))
	}

	c.JSON(http.StatusOK, response.CreateUserFromDomain(user))
}

// Create action for [post] /users/
func (u User) Create(c *gin.Context) {
	var userRequest request.UserCreation
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.NewError("bad request"))
		return
	}

	user := &domain.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	u.userRepository.Create(user)

	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("User %s was created.", user.Email)})
}
