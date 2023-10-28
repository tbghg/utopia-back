package v3

import "errors"

type ResponseWithData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseWithoutData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	SuccessResponse = ResponseWithoutData{
		Code: SuccessCode,
		Msg:  "success",
	}

	ActionTypeInvalidError = errors.New("操作类型无效") // 操作类型无效
	UserIDInvalidError     = errors.New("用户id无效") // 用户id无效
)

const (
	SuccessCode = 2000 // 成功
	ErrorCode   = 4000 // 失败
)
