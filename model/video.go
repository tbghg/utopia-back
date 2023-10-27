package model

import (
	"gorm.io/plugin/soft_delete"
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
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag;index:idx_author_del,priority:2;index:idx_ctime_del,priority:2"` // 软删除

	AuthorID uint   `gorm:"not null;index:idx_author_del,priority:1" json:"author_id"` // 作者id
	PlayUrl  string `gorm:"type:varchar(64);not null" json:"play_url"`                 // 视频播放地址
	CoverUrl string `gorm:"type:varchar(64);not null" json:"cover_url"`                // 视频封面
	Describe string `gorm:"type:varchar(64)" json:"describe"`                          // 描述
}
