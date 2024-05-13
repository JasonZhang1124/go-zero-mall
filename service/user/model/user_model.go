package model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name     string
	Gender   int64
	Mobile   string
	Password string
}
