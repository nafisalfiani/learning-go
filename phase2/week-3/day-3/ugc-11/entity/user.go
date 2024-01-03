package entity

type User struct {
	UserID        int     `json:"user_id" gorm:"primaryKey"`
	Username      string  `json:"username" validate:"required" gorm:"uniqueIndex"`
	Password      string  `json:"password" validate:"required"`
	DepositAmount float64 `json:"deposit_amount" validate:"required"`
}

type RegisterRequest struct {
	Username      string  `json:"username"`
	Password      string  `json:"password"`
	DepositAmount float64 `json:"deposit_amount"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	User    *User  `json:"user"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string  `json:"message"`
	Token   *string `json:"token"`
}
