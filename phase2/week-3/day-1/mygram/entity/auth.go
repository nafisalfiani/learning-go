package entity

import "time"

type RegisterRequest struct {
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Address     string    `json:"address"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken,omitempty"`
	Message     string `json:"message"`
}
