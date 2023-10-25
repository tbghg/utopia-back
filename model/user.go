package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;not null" json:"username"` // 账号
	Password string `gorm:"column:password;not null" json:"password"` // 密码（MD5）
	Nickname string `gorm:"column:name;not null" json:"nickname"`     // 昵称
	Salt     string `gorm:"column:salt;not null" json:"salt"`         // 密码盐
}
