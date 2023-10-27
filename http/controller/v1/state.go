package v1

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
)

const (
	SuccessCode = 2000 // 成功
	ErrorCode   = 4000 // 失败
)
