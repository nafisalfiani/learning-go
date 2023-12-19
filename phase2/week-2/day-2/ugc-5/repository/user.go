package repository

import (
	"data-center/entity"
	"database/sql"
)

type user struct {
	db *sql.DB
}

type Interface interface {
	List() ([]entity.User, error)
	Get(email string) (entity.User, error)
	Create(branch entity.User) (entity.User, error)
}

func InitUser(db *sql.DB) Interface {
	return &user{
		db: db,
	}
}

// List fetch a list of users
func (u *user) List() ([]entity.User, error) {
	var users []entity.User
	rows, err := u.db.Query("SELECT id, email, full_name, age, occupation, role FROM user")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.FullName,
			&user.Age,
			&user.Occupation,
			&user.Role,
		)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Get fetch specific user by given email
func (u *user) Get(email string) (entity.User, error) {
	var user entity.User
	err := u.db.QueryRow("SELECT id, email, full_name, age, occupation, role, password FROM user WHERE email = ?", email).Scan(
		&user.Id,
		&user.Email,
		&user.FullName,
		&user.Age,
		&user.Occupation,
		&user.Role,
		&user.Password,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Create creates new user based on given user struct
func (u *user) Create(user entity.User) (entity.User, error) {
	result, err := u.db.Exec("INSERT INTO user (email, full_name, age, occupation, role, password) VALUES (?, ?, ?, ?, ?, ?)",
		user.Email,
		user.FullName,
		user.Age,
		user.Occupation,
		user.Role,
		user.Password,
	)
	if err != nil {
		return user, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	if err := u.db.QueryRow("SELECT id, email, full_name, age, occupation, role FROM user WHERE id = ?", userID).Scan(
		&user.Id,
		&user.Email,
		&user.FullName,
		&user.Age,
		&user.Occupation,
		&user.Role,
	); err != nil {
		return user, err
	}

	return user, nil
}
