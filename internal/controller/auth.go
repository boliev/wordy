package controller

import (
	"github.com/boliev/wordy/internal/request"
	"github.com/boliev/wordy/internal/response"
	"github.com/boliev/wordy/internal/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Auth controller struct
type Auth struct {
	Authenticator *user.Authenticator
}

// CreateAuthController controller for /auth/
func CreateAuthController(authenticator *user.Authenticator) *Auth {
	return &Auth{
		Authenticator: authenticator,
	}
}

// Auth action for [post] /auth/
func (a Auth) Auth(c *gin.Context) {
	var AuthRequest request.UserAuth
	if err := c.ShouldBindJSON(&AuthRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.NewError("bad request"))
		return
	}

	auth, err := a.Authenticator.Auth(AuthRequest.Email, AuthRequest.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewError(err.Error()))
	}

	c.JSON(http.StatusOK, response.CreateAuthFromDomain(auth))
}
