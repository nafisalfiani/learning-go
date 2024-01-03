package repository

import (
	"avengers-commerce/entity"

	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

type UserInterface interface {
	List() ([]entity.User, error)
	Get(username string) (entity.User, error)
	Create(user entity.User) (entity.User, error)
}

// initUser create user repository
func initUser(db *gorm.DB) UserInterface {
	return &user{
		db: db,
	}
}

// List returns list of users
func (s *user) List() ([]entity.User, error) {
	users := []entity.User{}
	if err := s.db.Preload("Enrollments").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

// Get returns specific user by email
func (s *user) Get(username string) (entity.User, error) {
	user := entity.User{}
	if err := s.db.First(&user, &entity.User{Username: username}).Error; err != nil {
		return user, err
	}

	return user, nil
}

// Create creates new user data
func (s *user) Create(user entity.User) (entity.User, error) {
	if err := s.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
