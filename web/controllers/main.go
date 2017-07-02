package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MainController struct {
	beego.Controller
	BaseController
}

func (ctx *MainController) Prepare() {
	v := ctx.GetSession("admin")
	fmt.Println(v)
	if v == nil {
		ctx.Redirect("/admin/login", 302)
	}
}

// @router / [get]
func (c *MainController) Index() {
	c.Data["Name"] = "main"
	c.Data["model"] = "login"
	c.TplName = "index.html"
}

// @router /main/export [get]
func (c *MainController) Export() {
	list := GetTaskData()
	fmt.Println(list)
	c.Ctx.Output.Header("Content-type", "application/octet-stream")
	c.Ctx.Output.Header("Content-Disposition", "attachment; filename=\"hello.txt\"")
	c.Ctx.WriteString("hello")
}

type TaskInfo struct {
	Id     int
	Is_del int
}

func GetTaskData() []TaskInfo {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/test")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("like")
	results := make([]TaskInfo, 2)
	//err = c.Find(bson.M{"name": "jh"}).All(&results)
	err = c.Find(bson.M{}).All(&results)
	if err != nil {
		fmt.Println(err)
	}
	return results
}
