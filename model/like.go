package model

import (
	"time"
)

// Like 点赞表
//
// 视频ID-用户ID	唯一索引
// 视频ID-Status	索引
// 更新时间 		索引
type Like struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"index"`
	Status    bool      `gorm:"not nulls;index:idx_vid_status,priority:2"`                                                   // 状态 true:点赞 false:取消点赞
	VideoID   uint      `gorm:"not null;uniqueIndex:idx_uid_vid,priority:2;index:idx_vid_status,priority:1" json:"video_id"` // 视频id
	UserID    uint      `gorm:"not null;uniqueIndex:idx_uid_vid,priority:1" json:"user_id"`                                  // 用户id
}
