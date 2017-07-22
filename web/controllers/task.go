package controllers

import (
	"fmt"
	"web/models"

	"github.com/astaxie/beego"
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
	// select machines
	var mresults []models.MachineInfo
	machines := make(map[string]models.MachineInfo)
	mresults = models.GetMachines()
	for _, mresult := range mresults {
		machines[mresult.Id.Hex()] = mresult
	}

	// select tasks
	var results []models.TaskInfo
	results = models.GetTasks()

	var list []map[string]string

	for index, result := range results {
		item := make(map[string]string)
		item["dataid"] = result.Id.Hex()
		item["name"] = result.Name
		item["exectime"] = result.Exectime
		item["cmd"] = result.Cmd
		item["machineid"] = result.Machineid
		item["machine_name"] = ""
		if result.Machineid != "" {
			if _, ok := machines[result.Machineid]; ok {
				item["machine_name"] = machines[result.Machineid].Name
			}
		}

		list = append(list, item)
		fmt.Println(index, result.Name)
	}

	c.Data["list"] = list
	c.Data["machines"] = machines
	c.TplName = "task/tasks.html"
}

// @router /task/add [post]
func (ctx *TaskController) Add() {
	empty := []string{}
	var ret interface{}

	task_name := ctx.GetString("task_name")
	task_exectime := ctx.GetString("task_exectime")
	task_cmd := ctx.GetString("task_cmd")
	task_machineid := ctx.GetString("task_machineid")

	if "" == task_name {
		ret = ctx.ResponseJson(401, "task name is empty", empty)
	}
	if "" == task_exectime {
		ret = ctx.ResponseJson(401, "task exectime is empty", empty)
	}
	if "" == task_cmd {
		ret = ctx.ResponseJson(401, "task cmd is empty", empty)
	}
	if "" == task_machineid {
		ret = ctx.ResponseJson(401, "task machineid is empty", empty)
	}

	models.AddTask(task_name, task_exectime, task_cmd, task_machineid)

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
	task_machineid := ctx.GetString("task_machineid")

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
	if "" == task_machineid {
		ret = ctx.ResponseJson(401, "task machineid is empty", empty)
	}

	models.EditTask(dataid, task_name, task_exectime, task_cmd, task_machineid)

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
