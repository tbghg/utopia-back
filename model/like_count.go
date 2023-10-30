package model

import "time"

type LikeCount struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"index"`
	Count     int64     `gorm:"not null;default:0" json:"count"`                         // 点赞数
	VideoID   uint      `gorm:"not null;uniqueIndex:idx_vid,priority:2" json:"video_id"` // 视频id
}
