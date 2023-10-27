package utils

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"os"
	"testing"
)

func TestUploadFile(t *testing.T) {
	type args struct {
		localFile string
		key       string
		bucket    string
		secretKey string
		accessKey string
		mac       *qbox.Mac
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "upload file",
			args: args{
				localFile: "../../output/avatar.png",
				key:       "avatar.png",
				bucket:    os.Getenv("QINIU_BUCKET"),
				accessKey: os.Getenv("QINIU_ACCESS_KEY"),
				secretKey: os.Getenv("QINIU_SECRET_KEY"),
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.args.mac = qbox.NewMac(tt.args.accessKey, tt.args.secretKey)
		t.Run(tt.name, func(t *testing.T) {
			got, err := uploadFile(tt.args.localFile, tt.args.key, tt.args.bucket, tt.args.mac)
			if (err != nil) != tt.wantErr {
				t.Errorf("uploadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
