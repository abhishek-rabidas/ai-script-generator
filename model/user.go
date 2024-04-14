package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Uid       string `gorm:"not null; unique; size:64"`
	Profile   Profile
	ProfileId uint
	Scripts   []Script
}

type Profile struct {
	gorm.Model
	Uid   string `gorm:"not null; unique; size:64"`
	Wps   int
	Name  string `gorm:"not null; size:64"`
	Email string `gorm:"not null; size:64"`
}
