package routes

import (
	"github.com/kataras/iris/mvc"
	"github.com/yz124/superstar/bootstrap"
	"github.com/yz124/superstar/services"
	"github.com/yz124/superstar/web/controllers"
	"github.com/yz124/superstar/web/middleware"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	superstarService := services.NewSuperstarService()
	//注册路由，b.Party()实际调用的是*APIBuilder.Party()方法
	index := mvc.New(b.Party("/"))
	//注册服务处理
	index.Register(superstarService)
	//路由处理逻辑
	index.Handle(new(controllers.IndexController))
	//注册管理员路由，b.Party()实际调用的是*APIBuilder.Party()方法
	admin := mvc.New(b.Party("/admin"))
	//应用基本认证中间件
	admin.Router.Use(middleware.BasicAuth)
	//注册服务处理
	admin.Register(superstarService)
	//路由处理逻辑
	admin.Handle(new(controllers.AdminController))

	//b.Get("/follower/{id:long}", GetFollowerHandler)
	//b.Get("/following/{id:long}", GetFollowingHandler)
	//b.Get("/like/{id:long}", GetLikeHandler)
}

func Configuree(b *bootstrap.Bootstrapper) {
	superstarService := services.NewSuperstarService()
	index := mvc.New(b.Party("/"))
	index.Register(superstarService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	//在路由上面加中间件
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(superstarService)
	admin.Handle(new(controllers.AdminController))

}

