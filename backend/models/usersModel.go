package models

import "time"

type UserRole string

const (
	Normal UserRole = "Normal"
	Admin  UserRole = "Admin"
)

func (u UserRole) String() string {
	switch u {
	case Normal:
		return "Normal"
	case Admin:
		return "Admin"
	default:
		return "Unknown"
	}
}

type User struct {
	ID            string         `gorm:"type:uuid;primary_key;" json:"id"`
	Name          string         `gorm:"not null" json:"name"`
	Username      string         `gorm:"unique;not null" json:"username"`
	Email         string         `gorm:"unique;not null" json:"email"`
	Image         string         `gorm:"not null" json:"image"`
	Password      string         `json:"password" gorm:"not null"`
	Role          UserRole       `json:"role" gorm:"type:role_enum; not null"`
	CreatedAt     time.Time      `gorm:"autoCreateTime; not null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime; not null" json:"updated_at"`
	Conversations []Conversation `gorm:"many2many:conversations_users" json:"conversations"`
}
