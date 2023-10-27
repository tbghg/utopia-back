package utils

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"path"
	"utopia-back/config"
)

const (
	callbackURL      = "127.0.0.1:8080/api/v1/video/upload/callback"
	callbackBody     = `{"key":"$(key)","author_id":"$(x:author_id)"}`
	callbackBodyType = "application/json"
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

// GetCallbackToken 获取带有回调的Token
func GetCallbackToken() (upToken string) {
	accessKey := config.V.GetString("qiniu.accessKey")
	secretKey := config.V.GetString("qiniu.secretKey")
	bucket := config.V.GetString("qiniu.bucket")
	putPolicy := &storage.PutPolicy{
		Scope:            bucket,
		CallbackURL:      callbackURL,
		CallbackBody:     callbackBody,
		CallbackBodyType: callbackBodyType,
	}
	return getUpToken(putPolicy, accessKey, secretKey)
}

func QuickUploadFile(localFile, key string) (string, error) {
	accessKey := config.V.GetString("qiniu.accessKey")
	secretKey := config.V.GetString("qiniu.secretKey")
	bucket := config.V.GetString("qiniu.bucket")
	ret, err := uploadFile(localFile, key, bucket, accessKey, secretKey)
	// 拼接返回完整的url
	apiPath := config.V.GetString("qiniu.kodoApi")
	return path.Join(apiPath, ret), err
}

// uploadFile 上传文件到七牛云
func uploadFile(localFile, key, bucket, accessKey, secretKey string) (string, error) {

	upToken := getUpToken(&storage.PutPolicy{Scope: bucket}, accessKey, secretKey)
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

func getUpToken(putPolicy *storage.PutPolicy, accessKey string, secretKey string) (upToken string) {
	mac := qbox.NewMac(accessKey, secretKey)
	return putPolicy.UploadToken(mac)
}
