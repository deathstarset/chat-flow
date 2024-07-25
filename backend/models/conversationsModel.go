package models

import (
	"time"
)

type Conversation struct {
	ID        string    `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime; not null" json:"created_at"`
	Users     []User    `gorm:"many2many:conversations_users" json:"users"`
}
