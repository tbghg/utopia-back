package model

import (
	"time"
)

// Favorite 收藏表
//
// 视频ID-用户ID	唯一索引
// 视频ID-Status	联合索引
type Favorite struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    bool `gorm:"not null;default:0;index:idx_vid_status,priority:2" json:"status"`                            // 状态 true:收藏 false:取消收藏
	VideoID   uint `gorm:"not null;uniqueIndex:idx_uid_vid,priority:2;index:idx_vid_status,priority:1" json:"video_id"` // 视频id
	UserID    uint `gorm:"not null;uniqueIndex:idx_uid_vid,priority:1" json:"user_id"`                                  // 用户id
}
