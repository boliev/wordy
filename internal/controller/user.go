package controller

import (
	"fmt"
	"github.com/boliev/wordy/internal/repository"
	"github.com/boliev/wordy/internal/request"
	"github.com/boliev/wordy/internal/response"
	"github.com/boliev/wordy/internal/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// User struct
type User struct {
	userRepository repository.User
	userCreator    *user.Creator
}

// CreateUserController controller for /users/
func CreateUserController(userRepository repository.User, creator *user.Creator) *User {
	return &User{
		userRepository: userRepository,
		userCreator:    creator,
	}
}

// List action for [get] /users/
func (u User) List(c *gin.Context) {
	c.JSON(http.StatusOK, response.CreateUsersListFromDomain(u.userRepository.FindAll()))
}

// One action for [get] /users/:id
func (u User) One(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	usr, err := u.userRepository.Find(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewError(err.Error()))
	}

	c.JSON(http.StatusOK, response.CreateUserFromDomain(usr))
}

// Create action for [post] /users/
func (u User) Create(c *gin.Context) {
	var userRequest request.UserCreation
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.NewError("bad request"))
		return
	}

	newUser, err := u.userCreator.Create(userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewError("bad request"))
	}

	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("User %s was created.", newUser.Email)})
}
