package model

import "time"

type Comment struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"index"`
	VideoID   uint      `gorm:"not null;index" json:"video_id"` // 视频id
	UserID    uint      `gorm:"not null;index" json:"user_id"`  // 用户id
	Content   string    `gorm:"not null;content"`
}

type CommentInfo struct {
	UpdatedAt time.Time `json:"-"`
	Content   string    `json:"content"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
}
