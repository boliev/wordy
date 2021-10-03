package repository

import "github.com/boliev/wordy/internal/domain"

// User Repository interface
type User interface {
	Create(user *domain.User) error
	Find(id int) (*domain.User, error)
	FindAll() []*domain.User
}
