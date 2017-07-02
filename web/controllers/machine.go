package controllers

import (
	"fmt"
	"web/models"

	"github.com/astaxie/beego"
)

type MachineController struct {
	beego.Controller
	BaseController
}

func (ctx *MachineController) Prepare() {
	v := ctx.GetSession("admin")
	fmt.Println(v)
	if v == nil {
		ctx.Redirect("/admin/login", 302)
	}
}

// @router /machine/index [get]
func (c *MachineController) Index() {
	var results []models.MachineInfo
	results = models.GetMachines()
	list := make([]map[string]string, 0, 1)

	for index, result := range results {
		item := make(map[string]string)
		item["dataid"] = result.Id.Hex()
		item["name"] = result.Name
		item["hostname"] = result.Hostname
		item["ip"] = result.Ip
		item["account"] = result.Account
		list = append(list, item)
		fmt.Println(index, result.Id.Hex(), result.Name)
	}

	c.Data["list"] = list
	c.TplName = "machine/machines.html"
}

// @router /machine/add [post]
func (ctx *MachineController) Add() {
	empty := []string{}
	var ret interface{}

	machine_name := ctx.GetString("machine_name")
	machine_hostname := ctx.GetString("machine_hostname")
	machine_ip := ctx.GetString("machine_ip")
	machine_account := ctx.GetString("machine_account")

	if "" == machine_name {
		ret = ctx.ResponseJson(401, "machine name is empty", empty)
	}
	if "" == machine_hostname {
		ret = ctx.ResponseJson(401, "machine hostname is empty", empty)
	}
	if "" == machine_ip {
		ret = ctx.ResponseJson(401, "machine ip is empty", empty)
	}
	if "" == machine_account {
		ret = ctx.ResponseJson(401, "machine account is empty", empty)
	}

	models.AddMachine(machine_name, machine_hostname, machine_ip, machine_account)

	ret = ctx.ResponseJson(200, "success", empty)
	ctx.Data["json"] = ret
	ctx.ServeJSON()
}

// @router /machine/edit [post]
func (ctx *MachineController) Edit() {
	empty := []string{}
	var ret interface{}

	dataid := ctx.GetString("dataid")
	machine_name := ctx.GetString("machine_name")
	machine_hostname := ctx.GetString("machine_hostname")
	machine_ip := ctx.GetString("machine_ip")
	machine_account := ctx.GetString("machine_account")

	if "" == dataid {
		ret = ctx.ResponseJson(401, "dataid is empty", empty)
	}
	if "" == machine_name {
		ret = ctx.ResponseJson(401, "machine name is empty", empty)
	}
	if "" == machine_hostname {
		ret = ctx.ResponseJson(401, "machine hostname is empty", empty)
	}
	if "" == machine_ip {
		ret = ctx.ResponseJson(401, "machine ip is empty", empty)
	}
	if "" == machine_account {
		ret = ctx.ResponseJson(401, "machine account is empty", empty)
	}

	models.EditMachine(dataid, machine_name, machine_hostname, machine_ip, machine_account)

	ret = ctx.ResponseJson(200, "success", empty)
	ctx.Data["json"] = ret
	ctx.ServeJSON()
}

// @router /machine/remove [post]
func (ctx *MachineController) Remove() {
	empty := []string{}
	var ret interface{}

	dataid := ctx.GetString("dataid")

	if "" == dataid {
		ret = ctx.ResponseJson(401, "dataid is empty", empty)
	}

	models.RemoveMachine(dataid)

	ret = ctx.ResponseJson(200, "success", empty)
	ctx.Data["json"] = ret
	ctx.ServeJSON()
}
