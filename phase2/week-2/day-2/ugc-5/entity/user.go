package entity

type User struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	FullName   string `json:"full_name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
	Role       string `json:"role"`
}

type RegisterRequest struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,gte=8"`
	FullName   string `json:"full_name" validate:"required,gte=6,lte=15"`
	Age        int    `json:"age" validate:"required,gte=17"`
	Occupation string `json:"occupation" validate:"required"`
	Role       string `json:"role" validate:"required,oneof=admin superadmin"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}
