package ctx

import (
	"big-genius/core/errors"
	"big-genius/core/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

type Context struct {
	iris.Context
}

type Ctx context.Context

func (r Ctx) My() {
	println("ok")
}

type reply struct {
	ErrNo int32       `json:"errno"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

//
//func NewCtx(app context.Application) Context {
//	ctx := context.NewContext(app)
//
//	//timeout := 10 * time.Second
//	//_, cancel := ctx2.WithTimeout(ctx, timeout)
//	////ctx = ctxWithTimeout
//	//defer cancel()
//
//	//ctx.WithTimeout(r.Context(), h.dt)
//
//	timeout := 10 * time.Second
//	stdCtx, cancel := ctx2.WithTimeout(ctx2.Background(), timeout)
//	defer cancel()
//
//	// 将标准库的context.Context设置到Iris框架的Request中
//	// Set the standard library context as a value in iris.Context
//	ctx.Values().Set("stdCtx", stdCtx)
//
//	// Use iris.Context for further processing
//	stdCtxFromIris := ctx.Values().Get("stdCtx").(context.Context)
//	n := Context{
//		Context: &stdCtxFromIris,
//	}
//
//	return n
//}

func (c *Context) CloneCtx() Context {
	ctx := c.Clone()
	n := Context{
		Context: ctx,
	}
	return n
}

// Reply ...
func (c *Context) Reply(obj interface{}, err *errors.Error, codes ...int32) {
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

	err2 := c.JSON(r)
	if err2 != nil {
		log.Logger.Error(err2)
	}
}
