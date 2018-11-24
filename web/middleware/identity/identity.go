package identity

import (
	"time"
	"github.com/kataras/iris"
	"github.com/yz124/superstar/bootstrap"
	"runtime"
	"strconv"
)

// New returns a new handler which adds some headers and view data
// describing the application, i.e the owner, the startup time.

//设置header的中间件
func New(b *bootstrap.Bootstrapper) iris.Handler {
	return func(ctx iris.Context) {
		// response headers
		ctx.Header("App-Name", b.AppName)
		ctx.Header("App-Owner", b.AppOwner)
		ctx.Header("App-Since", time.Since(b.AppSpawnDate).String())
		ctx.Header("Server", runtime.GOOS)
		ctx.Header("Server-Arch", runtime.GOARCH)
		ctx.Header("GoMaxProcs", strconv.Itoa(runtime.GOMAXPROCS(0)))
		ctx.Header("Version", runtime.Version())
		// view data if ctx.View or c.Tmpl = "$page.html" will be called next.
		ctx.ViewData("AppName", b.AppName)
		ctx.ViewData("AppOwner", b.AppOwner)
		//转移控制权
		ctx.Next()
	}
}

func Neww(b *bootstrap.Bootstrapper) iris.Handler {
	return func(ctx iris.Context) {
		ctx.Header("app-name", b.AppName)
		ctx.Header("app-owner", b.AppOwner)
		ctx.Header("app-since", time.Since(b.AppSpawnDate).String())
		ctx.Header("server",runtime.GOOS)
		ctx.Header("server-arch", runtime.GOARCH)
		ctx.Header("gomaxprocs",strconv.Itoa(runtime.GOMAXPROCS(0)))
		ctx.Header("version", runtime.Version())
		ctx.Header("app-name",b.AppName)
		ctx.Header("app-owner",b.AppOwner)
		ctx.Next()
	}
}

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {
	h := New(b)
	//全局方式配置中间件
	b.UseGlobal(h)
}

func Configuree(b *bootstrap.Bootstrapper) {
	h := Neww(b)
	b.UseGlobal(h)
}
