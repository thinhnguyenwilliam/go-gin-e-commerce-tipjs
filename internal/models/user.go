package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey"` // already included in gorm.Model
	UUID      uuid.UUID `gorm:"column:uuid;type:varchar(255);not null;uniqueIndex:idx_uuid"`
	Username  string    `gorm:"column:user_name;type:varchar(255);not null"`
	IsActive  bool      `gorm:"column:is_active;type:boolean;default:true"`
	Roles     []Role    `gorm:"many2many:go_user_role"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"` // hashed password
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // optional: for soft delete
}

// Custom table name
func (User) TableName() string {
	return "go_db_user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New()
	return
}
