package models

import (
	"github.com/jinzhu/gorm"
	"github.com/go-sql-driver/mysql"
	_ "database/sql"
)

type User struct {
	gorm.Model
	Login  string `json: "login"`
	FirstName *string `json: "first_name`
	LastName *string `json: "last_name"`
	email *string `json: "email"`
}

type Task struct {
	gorm.Model
	Title       string
	Description string
	Priority    int
	CompletedAt mysql.NullTime
	IsCompleted bool
}