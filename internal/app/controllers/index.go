package controllers

import (
	ctx "big-genius/core/context"
)

func Index(ctx ctx.Context) {

	ctx.Writef("%s", "hello world")
}
