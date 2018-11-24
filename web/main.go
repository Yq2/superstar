package main

import (
	"github.com/yz124/superstar/bootstrap"
	"github.com/yz124/superstar/web/middleware/identity"
	"github.com/yz124/superstar/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("球星库1", "YQ")
	app.Bootstrap()
	//配置中间件
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
