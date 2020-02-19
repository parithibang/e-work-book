package setup

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func filters() {
	beego.InsertFilter("*", beego.BeforeRouter, FilterMethod)
}

func FilterMethod(ctx *context.Context) {
	if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
		ctx.Request.Method = strings.ToUpper(ctx.Input.Query("_method"))
	}
}
