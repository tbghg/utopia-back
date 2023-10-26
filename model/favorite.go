package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// Favorite 点赞表
//
// 视频ID-用户ID	唯一索引
// 视频ID-Del	索引
type Favorite struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag;index:idx_vid_del,priority:2"` // 软删除

	VideoID uint `gorm:"not null;uniqueIndex:idx_uid_vid,priority:2;index:idx_vid_del,priority:1" json:"video_id"` // 视频id
	UserID  uint `gorm:"not null;uniqueIndex:idx_uid_vid,priority:1" json:"user_id"`                               // 用户id
}
