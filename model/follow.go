package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// Follow 关注表
//
// 用户ID-粉丝ID-Del 唯一索引
// 粉丝ID-用户ID-Del 唯一索引
type Follow struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag;uniqueIndex:idx_uid_funId_del,priority:3;uniqueIndex:idx_funId_uid_del,priority:3"` // 软删除

	FunID  uint `gorm:"not null;uniqueIndex:idx_uid_funId_del,priority:2;uniqueIndex:idx_funId_uid_del,priority:1" json:"fun_id"`  // 关注的对象id
	UserID uint `gorm:"not null;uniqueIndex:idx_uid_funId_del,priority:1;uniqueIndex:idx_funId_uid_del,priority:2" json:"user_id"` // 用户id
}
