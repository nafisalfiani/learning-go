package repository

import (
	"mygram/entity"

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

func (u *user) List() ([]entity.User, error) {
	users := []entity.User{}
	if err := u.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u *user) Get(username string) (entity.User, error) {
	user := entity.User{}
	if err := u.db.First(&user, "username = ?", username).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *user) Create(user entity.User) (entity.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
