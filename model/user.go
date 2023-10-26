package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// User 用户表
//
// 用户名-Del 唯一索引
type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag;uniqueIndex:idx_uname_del,priority:2"` // 软删除

	Username string `gorm:"type:varchar(64);not null;uniqueIndex:idx_uname_del,priority:1" json:"username"` // 账号
	Password string `gorm:"type:varchar(64);not null" json:"password"`                                      // 密码（MD5）
	Nickname string `gorm:"type:varchar(64);not null" json:"nickname"`                                      // 昵称
	Avatar   string `gorm:"type:varchar(64);not null" json:"avatar"`                                        // 头像
	Salt     string `gorm:"type:varchar(8);not null" json:"salt"`                                           // 密码盐
}
