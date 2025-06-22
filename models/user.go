package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username string    `gorm:"size:20;unique;not null"`
	Email    string    `gorm:"size:30;unique;not null"`
	Password string    `gorm:"size:255;not null"`
	Role     string    `gorm:"size:15;not null;default:'user'"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
