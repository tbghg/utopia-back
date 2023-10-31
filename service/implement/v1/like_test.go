package v1

import (
	"testing"
	"utopia-back/cache"
	"utopia-back/database/abstract"
	"utopia-back/database/implement"
)

func TestLikeService_GetLikeCount(t *testing.T) {
	// 初始化数据库
	err := implement.TestInit()
	if err != nil {
		t.Errorf("TestInit() error = %v", err)
	}
	// 初始化redis
	err = cache.TestInit()

	type fields struct {
		LikeDal  abstract.LikeDal
		VideoDal abstract.VideoDal
	}
	type args struct {
		videoId uint
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantCount int64
		wantErr   bool
	}{
		{
			name: "test1",
			fields: fields{
				LikeDal:  &implement.LikeDal{},
				VideoDal: &implement.VideoDal{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LikeService{
				LikeDal:  tt.fields.LikeDal,
				VideoDal: tt.fields.VideoDal,
			}
			gotCount, err := l.GetLikeCount(tt.args.videoId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLikeCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.wantCount {
				t.Errorf("GetLikeCount() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
