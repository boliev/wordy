package memory

import (
	"fmt"
	"github.com/boliev/wordy/internal/domain"
)

// MemUserRepository User Repository memory implementation
type MemUserRepository struct {
	users     []*domain.User
	currentID int
}

// CreateMemUserRepository MemUserRepository constructor
func CreateMemUserRepository() *MemUserRepository {
	return &MemUserRepository{
		currentID: 0,
	}
}

// Create creates a user
func (m *MemUserRepository) Create(user *domain.User) {
	m.currentID++
	user.ID = m.currentID
	m.users = append(m.users, user)
}

// Find user by id
func (m *MemUserRepository) Find(id int) (*domain.User, error) {
	for _, user := range m.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, fmt.Errorf("no such user")
}

// FindAll get all users
func (m *MemUserRepository) FindAll() []*domain.User {
	return m.users
}
