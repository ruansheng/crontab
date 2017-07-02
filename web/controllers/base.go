package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (ctx *BaseController) ResponseJson(en int, em string, data interface{}) interface{} {
	ret := make(map[string]interface{})
	ret["en"] = en
	ret["em"] = em
	ret["data"] = data
	return ret
}
