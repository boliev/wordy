package middleware

import (
	"net/http"
	"strings"

	"github.com/boliev/wordy/internal/user"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	JwtParser user.JwtParser
}

type authHeader struct {
	Token string `header:"Authorization"`
}

func CreateAuthHandler(JwtParser user.JwtParser) *AuthHandler {
	return &AuthHandler{
		JwtParser: JwtParser,
	}
}

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

	userId, err := a.JwtParser.Parse(idTokenHeader[1])

	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
	}

	c.Set("userId", userId)
}
