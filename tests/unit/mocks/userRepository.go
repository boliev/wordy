package mocks

import (
	"github.com/boliev/wordy/internal/domain"
	"github.com/boliev/wordy/internal/repository"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	repository.User
	mock.Mock
}

func (u *UserRepositoryMock) FindByEmail(email string) (*domain.User, error) {
	args := u.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.User), args.Error(1)
}
