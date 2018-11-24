package controllers

import (
	"log"
	"time"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/yz124/superstar/models"
	"github.com/yz124/superstar/services"
)

type AdminController struct {
	Ctx     iris.Context
	Service services.SuperstarService //数据库操作service ---> dao
}

type AdminControllerr struct {
	Ctx iris.Context
	Service services.SuperstarService
}

//mvc.Result是一个接口，mvc.View实现了这个接口
func (c *AdminController) Get() mvc.Result {
	datalist := c.Service.GetAll()
	// set the model and render the view template.
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Datalist": datalist,
		},
		Layout: "admin/layout.html", // 不要跟前端的layout混用
	}
}

func (c *AdminControllerr) Get() mvc.Result {
	dataList := c.Service.GetAll()
	return mvc.View{
		Name:"admin/index.html",
		Data:iris.Map{
			"Title":"管理后台",
			"DataList": dataList,
		},
		Layout:"admin/layout.html",
	}
}

func (c *AdminController) GetEdit() mvc.Result {
	//取出参数url中的参数
	id, err := c.Ctx.URLParamInt("id")
	var data *models.StarInfo
	if err == nil {
		data = c.Service.Get(id)
	} else {
		data = &models.StarInfo{
			Id: 0,
		}
	}
	//fmt.Println(id, data)
	// set the model and render the view template.
	return mvc.View{
		Name: "admin/edit.html",
		Data: iris.Map{
			"Title": "管理后台",
			"info":  data,
		},
		Layout: "admin/layout.html", // 不要跟前端的layout混用
	}
}

func (c *AdminControllerr) GetEdit() mvc.Result {
	id ,err := c.Ctx.URLParamInt("id")
	var data *models.StarInfo
	if err != nil {
		data = c.Service.Get(id)
	} else {
		data = &models.StarInfo{Id:0}
	}
	//mvc.View用来处理get请求
	return mvc.View{
		Name:"admin/edit.html",
		Data:iris.Map{
			"Title":"管理后台",
			"Info":data,
		},
		Layout:"admin/layout.html",
	}
}

//mvc.Response也实现了mvc.Result接口
func (c *AdminController) PostSave() mvc.Result {
	info := models.StarInfo{}
	//将form表单的参数映射到model
	err := c.Ctx.ReadForm(&info)
	//fmt.Printf("%v\n", info)
	if err != nil {
		log.Fatal(err)
	}
	//更新记录
	if info.Id > 0 {
		info.SysUpdated = int(time.Now().Unix()) //更新时间
		//更新指定的列
		c.Service.Update(&info, []string{"name_zh", "name_en", "avatar",
			"birthday", "height", "weight", "club", "jersy", "coutry",
			"birthaddress", "feature", "moreinfo", "sys_updated"})
	} else {
		//创建新记录
		info.SysCreated = int(time.Now().Unix())
		c.Service.Create(&info)
	}
	//mvc.Response响应
	//mv.Response用来处理post请求
	return mvc.Response{
		Path: "/admin/",
	}
}

func (c *AdminControllerr) PostSave() mvc.Result {
	info := models.StarInfo{}
	err := c.Ctx.ReadForm(&info)
	if err != nil {
		log.Fatal(err)
	}
	if info.Id > 0 {
		info.SysUpdated = int(time.Now().Unix())
		c.Service.Update(&info, []string{
			"name_zh","name_en","avatar","birthday","height","weight","club","jersy","country","birthaddress",
		})
	} else {
		info.SysCreated = int(time.Now().Unix())
		c.Service.Create(&info)
	}
	//mv.Response用来处理post请求
	return mvc.Response{
		Path:"/admin/",
	}
}

func (c *AdminController) GetDelete() mvc.Result {
	//从URL里面读取id参数
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.Service.Delete(id)
	}
	//mvc.Response响应
	return mvc.Response{
		Path: "/admin/",
	}
}

func (c *AdminControllerr) GetDelete() mvc.Result {
	id ,err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.Service.Delete(id)
	}
	return mvc.Response{
		Path:"/admin/",
	}
}
