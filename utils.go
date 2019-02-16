package utils

import (
	"github.com/astaxie/beego/context"
)

func SetHTTPHeader(Ctx *context.Context) {
	Ctx.Output.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
	Ctx.Output.Header("Pragma", "no-cache")
	Ctx.Output.Header("Expires", "0")
	Ctx.Output.Header("X-Content-Type-Options", "nosniff")
	Ctx.Output.Header("Strict-Transport-Security", "max-age=31536000 ; includeSubDomains")
	Ctx.Output.Header("X-Frame-Options", "SAMEORIGIN")
	Ctx.Output.Header("X-XSS-Protection", "1; mode=block")
	Ctx.Output.Header("X-Content-Security-Policy", "default-src 'self'")
	//Ctx.Output.Header("X-WebKit-CSP", "default-src 'self'")
}
