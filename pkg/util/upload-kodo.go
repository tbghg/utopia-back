package utils

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"path"
	"utopia-back/config"
)

func buildCfg() (cfg storage.Config) {

	cfg = storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadongZheJiang2
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	return cfg
}

// UploadFile 上传文件到七牛云
func UploadFile(localFile, key, bucket, accessKey, secretKey string) (string, error) {

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	// 构建配置
	cfg := buildCfg()

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, nil)
	if err != nil {
		return "", err
	}
	return ret.Key, nil
}

func QuickUploadFile(localFile, key string) (string, error) {
	assessKey := config.V.GetString("qiniu.accessKey")
	secretKey := config.V.GetString("qiniu.secretKey")
	bucket := config.V.GetString("qiniu.bucket")
	ret, err := UploadFile(localFile, key, bucket, assessKey, secretKey)
	// 拼接返回完整的url
	apiPath := config.V.GetString("qiniu.kodoApi")
	return path.Join(apiPath, ret), err
}
