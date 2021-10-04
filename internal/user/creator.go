package user

import (
	"github.com/boliev/wordy/internal/domain"
	"github.com/boliev/wordy/internal/repository"
	"github.com/boliev/wordy/internal/request"
)

// Creator creates a user from request
type Creator struct {
	userRepository repository.User
}

// CreateUserCreator Creator constructor
func CreateUserCreator(userRepository repository.User) *Creator {
	return &Creator{
		userRepository: userRepository,
	}
}

// Create a user
func (c Creator) Create(request request.UserCreation) (*domain.User, error) {
	hashedPassword, err := hashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:    request.Email,
		Password: string(hashedPassword),
	}
	err = c.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
