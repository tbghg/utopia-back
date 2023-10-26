package v1

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SuccessCode = 2000 // 成功
	ErrorCode   = 4000 // 失败
)
