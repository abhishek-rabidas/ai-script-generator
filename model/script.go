package model

import "gorm.io/gorm"

type Script struct {
	gorm.Model
	Uid      string `gorm:"not null; unique; size:64"`
	Account  Account
	Platform string
	Duration int
	Topic    string
	Content  string `gorm:"size:10240"`
}
