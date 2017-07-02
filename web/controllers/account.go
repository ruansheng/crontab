package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"web/models"
)

type AccountController struct {
	beego.Controller
	BaseController
}

func (ctx *AccountController) Prepare() {
	v := ctx.GetSession("admin")
	fmt.Println(v)
	if v == nil {
		ctx.Redirect("/admin/login", 302)
	}
}

// @router /account/index [get]
func (c *AccountController) Index() {
	var results []models.AccountInfo
	results = models.GetAccounts()

	list := make([]map[string]string, 0, 1)

	for index, result := range results {
		item := make(map[string]string)
		item["dataid"] = result.Id.Hex()
		item["name"] = result.Name
		item["phonenumber"] = result.Phonenumber
		item["email"] = result.Email
		list = append(list, item)
		fmt.Println(index, result.Name)
	}

	c.Data["list"] = list
	c.TplName = "account/accounts.html"
}

// @router /account/add [post]
func (ctx *AccountController) Add() {
	empty := []string{}
	var ret interface{}

	name := ctx.GetString("name")
	password := ctx.GetString("password")
	phonenumber := ctx.GetString("phonenumber")
	email := ctx.GetString("email")

	if "" == name {
		ret = ctx.ResponseJson(401, "name is empty", empty)
	}
	if "" == password {
		ret = ctx.ResponseJson(401, "password is empty", empty)
	}
	if "" == phonenumber {
		ret = ctx.ResponseJson(401, "phonenumber is empty", empty)
	}
	if "" == email {
		ret = ctx.ResponseJson(401, "email is empty", empty)
	}

	models.AddAccount(name, password, phonenumber, email)

	ret = ctx.ResponseJson(200, "success", empty)
	ctx.Data["json"] = ret
	ctx.ServeJSON()
}

// @router /account/edit [post]
func (ctx *AccountController) Edit() {
	empty := []string{}
	var ret interface{}

	dataid := ctx.GetString("dataid")
	name := ctx.GetString("name")
	password := ctx.GetString("password")
	phonenumber := ctx.GetString("phonenumber")
	email := ctx.GetString("email")

	if "" == dataid {
		ret = ctx.ResponseJson(401, "dataid name is empty", empty)
	}
	if "" == name {
		ret = ctx.ResponseJson(401, "name is empty", empty)
	}
	if "" == phonenumber {
		ret = ctx.ResponseJson(401, "phonenumber is empty", empty)
	}
	if "" == email {
		ret = ctx.ResponseJson(401, "email is empty", empty)
	}

	models.EditAccount(dataid, name, password, phonenumber, email)

	ret = ctx.ResponseJson(200, "success", empty)
	ctx.Data["json"] = ret
	ctx.ServeJSON()
}

// @router /account/remove [post]
func (ctx *AccountController) Remove() {
	empty := []string{}
	var ret interface{}

	dataid := ctx.GetString("dataid")
	if "" == dataid {
		ret = ctx.ResponseJson(401, "dataid is empty", empty)
	}

	models.RemoveAccount(dataid)

	ret = ctx.ResponseJson(200, "success", empty)
	ctx.Data["json"] = ret
	ctx.ServeJSON()
}
