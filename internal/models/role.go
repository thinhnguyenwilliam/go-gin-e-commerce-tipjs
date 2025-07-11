package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        int64  `gorm:"column:id;type:int;not null;primaryKey;autoIncrement;comment:'Primary key is ID'"`
	RoleName  string `gorm:"column:role_name;type:varchar(255);not null"`
	RoleNote  string `gorm:"column:role_note;type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Optional soft delete
}

func (r *Role) TableName() string {
	return "go_db_role"
}
