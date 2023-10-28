package model

import (
	"time"
)

// Follow 关注表
//
// 用户ID-粉丝ID 唯一索引
// 粉丝ID-用户ID 唯一索引
type Follow struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    bool `gorm:"not null;uniqueIndex:idx_uid_funId_status,priority:3;uniqueIndex:idx_funId_uid_status,priority:3" json:"status"`  // 状态 true:关注 false:取消关注
	FunID     uint `gorm:"not null;uniqueIndex:idx_uid_funId_status,priority:2;uniqueIndex:idx_funId_uid_status,priority:1" json:"fun_id"`  // 关注的对象id
	UserID    uint `gorm:"not null;uniqueIndex:idx_uid_funId_status,priority:1;uniqueIndex:idx_funId_uid_status,priority:2" json:"user_id"` // 用户id
}
