package user

import (
	"fmt"
	"github.com/boliev/wordy/internal/domain"
	"github.com/boliev/wordy/internal/repository"
)

// Authenticator struct
type Authenticator struct {
	userRepository repository.User
	jwtService     *JwtService
}

// CreateUserAuthenticator Authenticator constructor
func CreateUserAuthenticator(userRepository repository.User, jwtService *JwtService) *Authenticator {
	return &Authenticator{
		userRepository: userRepository,
		jwtService:     jwtService,
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

	auth, err := a.jwtService.Create(int(usr.ID))
	if err != nil {
		return nil, err
	}

	return auth, nil
}
