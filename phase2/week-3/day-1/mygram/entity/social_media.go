package entity

type SocialMedia struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url" gorm:"column:social_media_url"`
	UserID         int    `json:"user_id"`
	User           User   `json:"user" gorm:"foreignKey:UserID"`
}

type SocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required"`
}
