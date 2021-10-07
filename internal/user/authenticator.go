package user

import (
	"fmt"
	"github.com/boliev/wordy/internal/domain"
	"github.com/boliev/wordy/internal/repository"
)

// Authenticator struct
type Authenticator struct {
	userRepository repository.User
	jwtCreator     JwtCreator
}

// CreateUserAuthenticator Authenticator constructor
func CreateUserAuthenticator(userRepository repository.User, jwtCreator JwtCreator) *Authenticator {
	return &Authenticator{
		userRepository: userRepository,
		jwtCreator:     jwtCreator,
	}
}

// Auth authenticates user
func (a Authenticator) Auth(email string, password string) (*domain.UserAuth, error) {
	usr, err := a.userRepository.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("no such user")
	}

	err = checkPassword(usr.Password, password)
	if err != nil {
		return nil, fmt.Errorf("no such user")
	}

	auth, err := a.jwtCreator.Create(int(usr.ID))
	if err != nil {
		return nil, err
	}

	return auth, nil
}
