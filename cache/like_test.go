package cache

import (
	"fmt"
	"testing"
	"utopia-back/pkg/logger"
)

func TestBuildUserLikedVideos(t *testing.T) {
	TestInit()
	logger.InitLogger(
		100,
		7,
		7,
		false,
		"debug",
	)
	type args struct {
		key    string
		vid    []uint
		minVid uint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				key:    "v3:like:1",
				vid:    []uint{100, 232, 354, 623},
				minVid: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BuildUserLikedVideos(tt.args.key, tt.args.vid, tt.args.minVid)
		})
	}
}

func Test_rebuildUserLikedVideos(t *testing.T) {
	TestInit()
	logger.InitLogger(
		100,
		7,
		7,
		false,
		"debug",
	)
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				key: "v3:like:1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rebuildUserLikedVideos(tt.args.key)
		})
	}
}

func TestIsUserLikedVideo(t *testing.T) {
	TestInit()
	logger.InitLogger(
		100,
		7,
		7,
		false,
		"debug",
	)
	type args struct {
		uid uint
		vid uint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "在缓存的数据",
			args: args{
				uid: 1,
				vid: 232,
			},
		},
		{
			name: "不存在的key",
			args: args{
				uid: 2,
				vid: 232,
			},
		},
		{
			name: "不存在的冷数据",
			args: args{
				uid: 1,
				vid: 99,
			},
		},
		{
			name: "不存在的热数据",
			args: args{
				uid: 1,
				vid: 23432,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLiked, gotState := IsUserLikedVideo(tt.args.uid, tt.args.vid)
			fmt.Println(gotLiked, gotState)
		})
	}
}

func TestIsUserLikedVideos(t *testing.T) {
	TestInit()
	logger.InitLogger(
		100,
		7,
		7,
		false,
		"debug",
	)
	type args struct {
		uid      uint
		videoIds []uint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "存在的用户",
			args: args{
				uid:      1,
				videoIds: []uint{12, 234, 100, 232},
			},
		}, {
			name: "不在缓存的用户",
			args: args{
				uid:      2,
				videoIds: []uint{12, 234, 100, 232},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotState := IsUserLikedVideos(tt.args.uid, tt.args.videoIds)
			fmt.Println(gotResult, gotState)
		})
	}
}
