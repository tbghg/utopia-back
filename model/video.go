package model

import (
	"gorm.io/gorm"
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
	DeletedAt gorm.DeletedAt `gorm:"index:idx_author_del,priority:2;index:idx_ctime_del,priority:2"`

	AuthorID    uint   `gorm:"not null;index:idx_author_del,priority:1" json:"author_id"` // 作者id
	PlayUrl     string `gorm:"type:varchar(256);not null" json:"play_url"`                // 视频播放地址
	CoverUrl    string `gorm:"type:varchar(256);not null" json:"cover_url"`               // 视频封面
	VideoTypeID uint   `gorm:"not null" json:"video_type_id"`                             // 视频类型
	Describe    string `gorm:"type:varchar(256)" json:"describe"`                         // 描述
	Title       string `gorm:"type:varchar(256)" json:"title"`                            // 标题
}

// VideoInfo 视频信息
type VideoInfo struct {
	ID        uint      `json:"id"`         // 视频id
	CreatedAt time.Time `json:"created_at"` // 视频创建时间

	PlayUrl     string `json:"play_url"`      // 视频播放地址
	CoverUrl    string `json:"cover_url"`     // 视频封面
	VideoTypeID uint   `json:"video_type_id"` // 视频类型
	Describe    string `json:"describe"`      // 描述
	Title       string `json:"title"`         // 标题

	Author AuthorInfo `json:"author"` // 作者信息

	IsFollow   bool `json:"is_follow"`   // 是否关注该作者
	IsLike     bool `json:"is_like"`     // 是否点赞过
	IsFavorite bool `json:"is_favorite"` // 是否收藏过

	LikeCount     int `json:"like_count"`     // 点赞数
	FavoriteCount int `json:"favorite_count"` // 点赞数
}

// AuthorInfo 作者信息
type AuthorInfo struct {
	ID          uint   `json:"id"`           // 用户id
	Nickname    string `json:"nickname"`     // 昵称
	Avatar      string `json:"avatar"`       // 头像
	Username    string `json:"username"`     // 用户名
	FansCount   int64  `json:"fans_count"`   // 粉丝数
	FollowCount int64  `json:"follow_count"` // 关注数
	VideoCount  int64  `json:"video_count"`  // 视频数
}

type VideoCount struct {
	VideoID uint
	Count   int
	Score   float64
}

type CallbackItem struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Key  string `json:"key"`
}
