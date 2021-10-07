package user_test

import (
	"fmt"
	"github.com/boliev/wordy/internal/domain"
	"github.com/boliev/wordy/internal/user"
	"github.com/boliev/wordy/tests/unit/mocks"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestAuthSuccess(t *testing.T) {
	userRepository := new(mocks.UserRepositoryMock)
	jwtCreator := &mocks.JwtCreatorMock{}
	usr := getUser(1, "some@email.com", "123456")
	userAuth := getUserAuth("111222333", 1)
	userRepository.On("FindByEmail", "some@email.com").Return(usr, nil).Once()
	jwtCreator.On("Create", 1).Return(userAuth, nil).Once()
	service := user.CreateUserAuthenticator(userRepository, jwtCreator)

	resultUserAuth, err := service.Auth("some@email.com", "123456")

	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	assert.Equal(t, resultUserAuth.Token, "111222333")
}

func TestWrongPassword(t *testing.T) {
	userRepository := new(mocks.UserRepositoryMock)
	jwtCreator := &mocks.JwtCreatorMock{}
	usr := getUser(1, "some@email.com", "123457")
	userAuth := getUserAuth("111222333", 1)
	userRepository.On("FindByEmail", "some@email.com").Return(usr, nil).Once()
	jwtCreator.On("Create", 1).Return(userAuth, nil).Once()
	service := user.CreateUserAuthenticator(userRepository, jwtCreator)

	_, err := service.Auth("some@email.com", "123456")

	assert.NotNil(t, err)
	assert.Equal(t, "no such user", err.Error())
}

func TestNoSuchUser(t *testing.T) {
	userRepository := new(mocks.UserRepositoryMock)
	jwtCreator := &mocks.JwtCreatorMock{}
	userRepository.On("FindByEmail", "some@email.com").Return(nil, fmt.Errorf("some error")).Once()
	jwtCreator.On("Create", 1).Return(nil, nil).Once()
	service := user.CreateUserAuthenticator(userRepository, jwtCreator)

	_, err := service.Auth("some@email.com", "123456")

	assert.NotNil(t, err)
	assert.Equal(t, "no such user", err.Error())
}

func getUser(id uint, email string, password string) *domain.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	return &domain.User{
		Model: gorm.Model{
			ID: id,
		},
		Email:    email,
		Password: string(hashedPassword),
	}
}
func getUserAuth(token string, days int) *domain.UserAuth {
	return &domain.UserAuth{
		Token:     token,
		ExpiresAt: time.Now().AddDate(0, 0, days),
	}
}
