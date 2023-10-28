package utils

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"path"
	"utopia-back/config"
)

const (
	callbackPath     = "/api/v1/video/upload/callback"
	callbackBody     = `{"key":"$(key)","fileType":"$(x:fileType)","uid":"$(x:uid)","cover_url":"$(x:cover_url)","describe":"$(x:describe)","video_type_id":"$(x:video_type_id)"}`
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
	bucket := config.V.GetString("qiniu.bucket")
	callbackUrl := config.V.GetString("server.ip") + config.V.GetString("server.port") + callbackPath
	putPolicy := &storage.PutPolicy{
		Scope:            bucket,
		CallbackURL:      callbackUrl,
		CallbackBody:     callbackBody,
		CallbackBodyType: callbackBodyType,
	}
	return putPolicy.UploadToken(GetMac())
}

func QuickUploadFile(localFile, key string) (string, error) {
	bucket := config.V.GetString("qiniu.bucket")
	ret, err := uploadFile(localFile, key, bucket, GetMac())
	// 拼接返回完整的url
	apiPath := config.V.GetString("qiniu.kodoApi")
	return path.Join(apiPath, ret), err
}

// uploadFile 上传文件到七牛云
func uploadFile(localFile, key, bucket string, mac *qbox.Mac) (string, error) {

	putPolicy := &storage.PutPolicy{Scope: bucket}
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

func GetMac() *qbox.Mac {
	accessKey := config.V.GetString("qiniu.accessKey")
	secretKey := config.V.GetString("qiniu.secretKey")
	return qbox.NewMac(accessKey, secretKey)
}
