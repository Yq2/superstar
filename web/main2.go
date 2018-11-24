package main

import (
	"github.com/yz124/superstar/bootstrap"
	"github.com/yz124/superstar/web/middleware/identity"
	"github.com/yz124/superstar/web/routes"
)

func newApp2() *bootstrap.Bootstrapper {
	app := bootstrap.New("球星库2", "YQ")
	app.Bootstrap()
	//配置中间件
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp2()
	app.Listen(":8081")
}
