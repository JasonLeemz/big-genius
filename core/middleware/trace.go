package trace

import (
	"github.com/kataras/iris/v12"
)

// Inject TODO ...
func Inject(ctx iris.Context) {
	shareInformation := "this is a sharable information between handlers"

	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	ctx.Values().Set("info", shareInformation)

	ctx.Next()
}
