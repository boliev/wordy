package user

import "github.com/boliev/wordy/internal/domain"

// JwtCreator interface
type JwtCreator interface {
	Create(id int) (*domain.UserAuth, error)
}
