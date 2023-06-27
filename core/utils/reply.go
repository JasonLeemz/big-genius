package utils

import (
	"big-genius/core/errors"
	"big-genius/core/log"
	"github.com/kataras/iris/v12/context"
)

type reply struct {
	ErrNo int32       `json:"errno"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

// Reply ...
func Reply(ctx *context.Context, obj interface{}, err *errors.Error, codes ...int32) {
	r := &reply{
		Data: obj,
	}

	code := int32(0)
	if len(codes) != 0 {
		code = codes[0]
		r.ErrNo = code
	}

	if err != nil {
		r.Msg = err.Error()
	} else {
		r.Msg = errors.GetErrMsg(code)
	}

	err2 := ctx.JSON(r)
	if err2 != nil {
		log.Logger.Error(err2)
	}
}
