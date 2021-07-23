package errorx

import "net/http"

const CommonErrCode = 1

type RespError struct {
	Status int
	Code   int
	Msg    string
}

func NewRespError(status int, code int, msg string) error {
	return &RespError{Status: status, Code: code, Msg: msg}
}

func NewCommonError(msg string) error {
	return NewRespError(http.StatusBadRequest, CommonErrCode, msg)
}

func (e *RespError) Error() string {
	return e.Msg
}

func (e *RespError) Data() interface{} {
	return map[string]interface{}{
		"code": e.Code,
		"msg":  e.Msg,
	}
}
