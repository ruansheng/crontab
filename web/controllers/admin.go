package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
	BaseController
}

func (ctx *AdminController) Prepare() {
	v := ctx.GetSession("admin")
	fmt.Println(v)
	if v != nil {
		ctx.Redirect("/", 302)
	}
}

// @router /admin/login [get]
func (c *AdminController) Login() {
	c.Data["Name"] = "crontab"
	c.Data["model"] = "login"
	c.TplName = "login.html"
}

// @router /admin/doLogin [post]
func (ctx *AdminController) DoLogin() {
	empty := []string{}
	var ret interface{}

	username := ctx.GetString("username")
	password := ctx.GetString("password")

	if username != "ruansheng" && password != "123" {
		ret = ctx.ResponseJson(401, "machine name is empty", empty)
		ctx.Data["json"] = ret
		ctx.ServeJSON()
	}

	// 设置session
	ctx.SetSession("admin", "ruansheng")
	ctx.Data["admin"] = "ruansheng"

	ret = ctx.ResponseJson(200, "success", empty)
	ctx.Data["json"] = ret
	ctx.ServeJSON()
}

// @router /admin/logout [get]
func (ctx *AdminController) Logout() {
	ctx.DelSession("admin")
	ctx.Data["admin"] = nil
	fmt.Println("admin->logout")
	ctx.Redirect("/admin/login", 302)
}
