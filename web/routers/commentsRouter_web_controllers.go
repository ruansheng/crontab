package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["web/controllers:AdminController"] = append(beego.GlobalControllerRouter["web/controllers:AdminController"],
		beego.ControllerComments{
			"Login",
			`/admin/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["web/controllers:AdminController"] = append(beego.GlobalControllerRouter["web/controllers:AdminController"],
		beego.ControllerComments{
			"DoLogin",
			`/admin/doLogin`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["web/controllers:AdminController"] = append(beego.GlobalControllerRouter["web/controllers:AdminController"],
		beego.ControllerComments{
			"Logout",
			`/admin/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["web/controllers:MachineController"] = append(beego.GlobalControllerRouter["web/controllers:MachineController"],
		beego.ControllerComments{
			"Index",
			`/machine/index`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["web/controllers:MachineController"] = append(beego.GlobalControllerRouter["web/controllers:MachineController"],
		beego.ControllerComments{
			"Add",
			`/machine/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["web/controllers:MachineController"] = append(beego.GlobalControllerRouter["web/controllers:MachineController"],
		beego.ControllerComments{
			"Edit",
			`/machine/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["web/controllers:MachineController"] = append(beego.GlobalControllerRouter["web/controllers:MachineController"],
		beego.ControllerComments{
			"Remove",
			`/machine/remove`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["web/controllers:MainController"] = append(beego.GlobalControllerRouter["web/controllers:MainController"],
		beego.ControllerComments{
			"Index",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["web/controllers:MainController"] = append(beego.GlobalControllerRouter["web/controllers:MainController"],
		beego.ControllerComments{
			"Export",
			`/main/export`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["web/controllers:TaskController"] = append(beego.GlobalControllerRouter["web/controllers:TaskController"],
		beego.ControllerComments{
			"Index",
			`/task/index`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["web/controllers:TaskController"] = append(beego.GlobalControllerRouter["web/controllers:TaskController"],
		beego.ControllerComments{
			"Add",
			`/task/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["web/controllers:TaskController"] = append(beego.GlobalControllerRouter["web/controllers:TaskController"],
		beego.ControllerComments{
			"Edit",
			`/task/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["web/controllers:TaskController"] = append(beego.GlobalControllerRouter["web/controllers:TaskController"],
		beego.ControllerComments{
			"Remove",
			`/task/remove`,
			[]string{"post"},
			nil})

}
