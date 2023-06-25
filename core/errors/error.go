package errors

import (
	"big-genius/core/log"
	"encoding/json"
)

type Error struct {
	err error
	no  int32  `json:"no"`
	msg string `json:"msg"`
}

func (e *Error) ErrNo() int32 {
	return e.no
}

func (e *Error) ToString() string {
	s, err := json.Marshal(e)
	if err != nil {
		log.Logger.Error(err)
	}
	return string(s)
}

func (e *Error) SetErrNo(code int32) {
	e.no = code
}

func (e *Error) SetError(msg string) {
	e.msg = msg
}

func (e *Error) Error() string {
	errMsg := "success"

	if e.err != nil {
		errMsg = e.err.Error()
	}
	if e.msg != "" {
		errMsg += ": [" + e.msg + "]"
	}

	return errMsg
}

type Impl interface {
	Error() string
	ErrNo() int32
	ToString() string

	SetErrNo(code int32)
	SetError(msg string)
}

func GenErr(e error, codes ...int32) *Error {
	var err *Error
	code := int32(0)
	if len(codes) != 0 {
		code = codes[0]
	}

	if code != 0 && e != nil {
		err.SetErrNo(code)
	}

	// 返回自定义错误
	if code != 0 && e == nil {
		err = &Error{
			no: code,
		}

		if msg, ok := errorMessages[int(code)]; ok {
			err.SetError(msg)
		} else {
			err.SetError(errorMessages[-1])
		}
	}

	// 同时存在系统错误和自定义code, 错误码以自定义为准
	if code != 0 && e != nil {
		err = &Error{
			err: e,
			no:  code,
		}
		if msg, ok := errorMessages[int(code)]; ok {
			err.SetError(e.Error() + ": [" + msg + "]")
		} else {
			err.SetError(errorMessages[-1])
		}
	}

	// 自定义code=0, err不为空
	if code == 0 && e != nil {
		err = &Error{
			err: e,
			no:  -1,
		}
		err.SetError(e.Error())
	}
	return err
}

func GetErrMsg(code int32) string {
	if msg, ok := errorMessages[int(code)]; ok {
		return msg
	} else {
		return errorMessages[-1]
	}
}
