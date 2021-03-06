package bootstrap

import (
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"
	"time"
	"github.com/yz124/superstar/conf"
)

type Configurator func(*Bootstrapper)

// 使用Go内建的嵌入机制(匿名嵌入)，允许类型之前共享代码和数据
// （Bootstrapper继承和共享 iris.Application ）
// 参考文章： https://hackthology.com/golangzhong-de-mian-xiang-dui-xiang-ji-cheng.html
type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
	Sessions *sessions.Sessions
}

// New returns a new Bootstrapper.
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper {
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(), //返回一个application
	}
	for _, cfg := range cfgs {
		cfg(b)
	}
	return b
}

// SetupViews loads the templates.
func (b *Bootstrapper) SetupViews(viewsDir string) {
	//创建视图引擎
	htmlEngine := iris.HTML(viewsDir, ".html").Layout("shared/layout.html")
	// 每次重新加载模版（线上关闭它）
	htmlEngine.Reload(false)
	// 给模版内置各种定制的方法
	htmlEngine.AddFunc("FromUnixtimeShort", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeformShort) //时间格式化,年-月-日
	})
	htmlEngine.AddFunc("FromUnixtime", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeform) //年-月-日-时-分-秒
	})
	//*iris.Application 注册视图引擎
	b.RegisterView(htmlEngine)
}

// SetupSessions initializes the sessions, optionally.
//创建session
func (b *Bootstrapper) SetupSessions(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
	b.Sessions = sessions.New(sessions.Config{
		Cookie:   "SECRET_SESS_COOKIE_" + b.AppName,
		Expires:  expires, //session有效期
		Encoding: securecookie.New(cookieHashKey, cookieBlockKey), //创建cookie 编码
	})
}

// SetupWebsockets prepares the websocket server.
//websocket配置
func (b *Bootstrapper) SetupWebsockets(endpoint string, onConnection websocket.ConnectionFunc) {
	//初始化一个websocket
	ws := websocket.New(websocket.Config{})
	//处理连接
	ws.OnConnection(onConnection)
	//endpoint服务端点
	b.Get(endpoint, ws.Handler())
	//向客户端写入js文件
	b.Any("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})
}

// SetupErrorHandlers prepares the http error handlers
// `(context.StatusCodeNotSuccessful`,  which defaults to < 200 || >= 400 but you can change it).
//  错误处理
func (b *Bootstrapper) SetupErrorHandlers() {
	//  注册错误处理
	//可以处理任何错误响应码
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map {
			"app":     b.AppName, //APP名字
			"status":  ctx.GetStatusCode(), //状态码
			"message": ctx.Values().GetString("message"), //从响应value中取出message
		}
		//如果URL参数里面存在json这个字段
		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			//那么把上面准备的err以JSON格式响应给客户端
			ctx.JSON(err)
			return
		}
		//否则就通过视图方式响应错误
		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}

const (
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./public/"  //静态目录
	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "favicon.ico"  //fav图标
)

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

// Bootstrap prepares our application.
//
// Returns itself.
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	//设置视图目录
	b.SetupViews("./views")
	//session有效期
	b.SetupSessions(
		24*time.Hour,
		[]byte("the-big-and-secret-fash-key-here"),
		[]byte("lot-secret-of-characters-big-too"),
	)
	//错误处理
	b.SetupErrorHandlers()

	// static files
	b.Favicon(StaticAssets + Favicon)
	b.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	// middleware, after static files
	//中间件在静态文件后面
	//宕机恢复中间件
	b.Use(recover.New())
	//日志记录
	b.Use(logger.New())
	return b
}

// Listen starts the http server with the specified "addr".
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}