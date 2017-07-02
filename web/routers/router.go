package routers

import (
	"github.com/astaxie/beego"
	"web/controllers"
)

func init() {
	beego.Include(&controllers.MainController{})
	beego.Include(&controllers.AdminController{})
	beego.Include(&controllers.AccountController{})
	beego.Include(&controllers.TaskController{})
	beego.Include(&controllers.MachineController{})
}
