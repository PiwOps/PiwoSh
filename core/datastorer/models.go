package datastorer

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint64 `json:"id" gorm:"primaryKey; uniqueIndex; not null"`
}

func (User) TableName() string {
	return "users"
}

type CommandsConfig struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey"`
	Key   string `gorm:"unique;not null"`
	Value string `gorm:"not null"`
}

func (CommandsConfig) TableName() string {
	return "commands_config"
}