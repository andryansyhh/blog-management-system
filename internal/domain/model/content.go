package model

import "time"

type Content struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	UserID     int       `json:"user_id"`
	User       User      `gorm:"foreignKey:UserID" json:"user"`
	CategoryID int       `json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID" json:"category"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
