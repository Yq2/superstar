package controllers

import (
	"log"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/yz124/superstar/datasource"
	"github.com/yz124/superstar/models"
	"github.com/yz124/superstar/services"
)

type IndexController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

func (c *IndexController) Get() mvc.Result {
	//c.Service.GetAll()
	//return mvc.Response{
	//	Text:"ok\n",
	//}
	datalist := c.Service.GetAll()
	//用数据渲染视图
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "球星库",
			"Datalist": datalist,
		},
	}
}

// http://localhost:8080/{id}
//GetBy 用来处理restful api URL上面带的参数
func (c *IndexController) GetBy(id int) mvc.Result {
	if id < 1 {
		return mvc.Response{
			Path: "/",
		}
	}
	data := c.Service.Get(id)
	//用数据渲染视图
	return mvc.View{
		Name: "info.html",
		Data: iris.Map{
			"Title": "球星库",
			"info":  data,
		},
	}
}

// http://localhost:8080/search?country=巴西
func (c *IndexController) GetSearch() mvc.Result {
	//取出URL后面跟的?country=xxx
	country := c.Ctx.URLParam("country")
	datalist := c.Service.Search(country)
	// set the model and render the view template.
	//用数据渲染视图
	return mvc.View {
		Name: "index.html",
		Data: iris.Map {
			"Title":    "球星库",
			"Datalist": datalist,
		},
	}
}

// 集群多服务器的时候，才用得上这个接口
// 性能优化的时候才考虑，加上本机的SQL缓存
// http://localhost:8080/clearcache
func (c *IndexController) GetClearcache() mvc.Result {
	err := datasource.InstanceMaster().ClearCache(&models.StarInfo{})
	if err != nil {
		log.Fatal(err)
	}
	// set the model and render the view template.
	return mvc.Response{
		Text: "xorm缓存清除成功",
	}
}
