package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"web/models"
)

type TaskController struct {
	beego.Controller
	BaseController
}

func (ctx *TaskController) Prepare() {
	v := ctx.GetSession("admin")
	fmt.Println(v)
	if v == nil {
		ctx.Redirect("/admin/login", 302)
	}
}

// @router /task/index [get]
func (c *TaskController) Index() {
	var results []models.TaskInfo
	results = models.GetTasks()

	list := make([]map[string]string, 0, 1)

	for index, result := range results {
		item := make(map[string]string)
		item["dataid"] = result.Id.Hex()
		item["name"] = result.Name
		item["exectime"] = result.Exectime
		item["cmd"] = result.Cmd
		list = append(list, item)
		fmt.Println(index, result.Name)
	}

	c.Data["list"] = list
	c.TplName = "task/tasks.html"
}

// @router /task/add [post]
func (ctx *TaskController) Add() {
	empty := []string{}
	var ret interface{}

	task_name := ctx.GetString("task_name")
	task_exectime := ctx.GetString("task_exectime")
	task_cmd := ctx.GetString("task_cmd")

	if "" == task_name {
		ret = ctx.ResponseJson(401, "task name is empty", empty)
	}
	if "" == task_exectime {
		ret = ctx.ResponseJson(401, "task exectime is empty", empty)
	}
	if "" == task_cmd {
		ret = ctx.ResponseJson(401, "task cmd is empty", empty)
	}

	models.AddTask(task_name, task_exectime, task_cmd)

	ret = ctx.ResponseJson(200, "success", empty)
	ctx.Data["json"] = ret
	ctx.ServeJSON()
}

// @router /task/edit [post]
func (ctx *TaskController) Edit() {
	empty := []string{}
	var ret interface{}

	dataid := ctx.GetString("dataid")
	task_name := ctx.GetString("task_name")
	task_exectime := ctx.GetString("task_exectime")
	task_cmd := ctx.GetString("task_cmd")

	if "" == dataid {
		ret = ctx.ResponseJson(401, "dataid name is empty", empty)
	}
	if "" == task_name {
		ret = ctx.ResponseJson(401, "task name is empty", empty)
	}
	if "" == task_exectime {
		ret = ctx.ResponseJson(401, "task exectime is empty", empty)
	}
	if "" == task_cmd {
		ret = ctx.ResponseJson(401, "task cmd is empty", empty)
	}

	models.EditTask(dataid, task_name, task_exectime, task_cmd)

	ret = ctx.ResponseJson(200, "success", empty)
	ctx.Data["json"] = ret
	ctx.ServeJSON()
}

// @router /task/remove [post]
func (ctx *TaskController) Remove() {
	empty := []string{}
	var ret interface{}

	dataid := ctx.GetString("dataid")

	if "" == dataid {
		ret = ctx.ResponseJson(401, "dataid is empty", empty)
	}

	models.RemoveTask(dataid)

	ret = ctx.ResponseJson(200, "success", empty)
	ctx.Data["json"] = ret
	ctx.ServeJSON()
}
