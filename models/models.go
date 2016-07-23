package models

import (
	"github.com/jinzhu/gorm"
	"github.com/go-sql-driver/mysql"
)

type User struct {
	gorm.Model
	Login  string
}

type Task struct {
	gorm.Model
	Title       string
	Description string
	Priority    int
	CompletedAt mysql.NullTime
	IsCompleted bool
}