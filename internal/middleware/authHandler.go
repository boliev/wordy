package middleware

import (
	"net/http"
	"strings"

	"github.com/boliev/wordy/internal/user"
	"github.com/gin-gonic/gin"
)

// AuthHandler struct
type AuthHandler struct {
	JwtParser user.JwtParser
}

type authHeader struct {
	Token string `header:"Authorization"`
}

// CreateAuthHandler AuthHandler constructor
func CreateAuthHandler(JwtParser user.JwtParser) *AuthHandler {
	return &AuthHandler{
		JwtParser: JwtParser,
	}
}

// Handle function
func (a AuthHandler) Handle(c *gin.Context) {
	header := &authHeader{}
	err := c.ShouldBindHeader(&header)
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
	}

	if header.Token == "" {
		c.AbortWithStatus(http.StatusForbidden)
	}

	idTokenHeader := strings.Split(header.Token, "Bearer ")

	userID, err := a.JwtParser.Parse(idTokenHeader[1])

	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
	}

	c.Set("userID", userID)
}
