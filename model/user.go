package model

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

//    }
type Model struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type User struct {
	Model
	Name    string
	Password string
}
