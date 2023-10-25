package model

import "gorm.io/gorm"

type TestUser struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;default:'';comment:'姓名'"`
	Age  int    `gorm:"type:int(11);not null;default:0;comment:'年龄'"`
}
