package mocks

import (
	"github.com/boliev/wordy/internal/domain"
	"github.com/stretchr/testify/mock"
)

type JwtCreatorMock struct {
	mock.Mock
}

func (j *JwtCreatorMock) Create(id int) (*domain.UserAuth, error) {
	args := j.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.UserAuth), args.Error(1)
}
