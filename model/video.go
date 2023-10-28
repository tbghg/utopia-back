package model

import (
	"time"
)

// Video 视频表
//
// 作者ID-Del	索引
// 发布时间-Del	索引
type Video struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"index:idx_ctime_del,priority:1" json:"created_at"`
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index:idx_author_del,priority:2;index:idx_ctime_del,priority:2"`

	AuthorID    uint   `gorm:"not null;index:idx_author_del,priority:1" json:"author_id"` // 作者id
	PlayUrl     string `gorm:"type:varchar(64);not null" json:"play_url"`                 // 视频播放地址
	CoverUrl    string `gorm:"type:varchar(64);not null" json:"cover_url"`                // 视频封面
	VideoTypeID uint   `gorm:"not null" json:"video_type_id"`                             // 视频类型
	Describe    string `gorm:"type:varchar(64)" json:"describe"`                          // 描述
}
