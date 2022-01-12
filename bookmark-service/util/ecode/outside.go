package ecode

import "github.com/MenciusCheng/superman/util/ecode"

type WrapResp struct {
	Code int         `json:"dm_error"`
	Msg  string      `json:"error_msg"`
	Data interface{} `json:"data"`
}

func NewWrapResp(data interface{}, err error) WrapResp {
	e := ecode.Cause(err)
	return WrapResp{
		Code: e.Code(),
		Msg:  e.Message(),
		Data: data,
	}
}
