package psql

import (
	"github.com/boliev/wordy/internal/domain"
	"gorm.io/gorm"
)

// UserRepository User Repository PostgreSQL implementation
type UserRepository struct {
	db *gorm.DB
}

// CreateUserRepository PostgreSQL UserRepository constructor
func CreateUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create creates a user
func (r *UserRepository) Create(user *domain.User) error {
	result := r.db.Create(user)

	return result.Error
}

// Find user by id
func (r *UserRepository) Find(id int) (*domain.User, error) {
	var user domain.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// FindAll get all users
func (r *UserRepository) FindAll() []*domain.User {
	var users []*domain.User
	r.db.Find(&users)

	return users
}
