package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["web/controllers:AccountController"] = append(beego.GlobalControllerRouter["web/controllers:AccountController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/account/index`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:AccountController"] = append(beego.GlobalControllerRouter["web/controllers:AccountController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/account/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:AccountController"] = append(beego.GlobalControllerRouter["web/controllers:AccountController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/account/edit`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:AccountController"] = append(beego.GlobalControllerRouter["web/controllers:AccountController"],
		beego.ControllerComments{
			Method: "Remove",
			Router: `/account/remove`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:AdminController"] = append(beego.GlobalControllerRouter["web/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/admin/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:AdminController"] = append(beego.GlobalControllerRouter["web/controllers:AdminController"],
		beego.ControllerComments{
			Method: "DoLogin",
			Router: `/admin/doLogin`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:AdminController"] = append(beego.GlobalControllerRouter["web/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/admin/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:MachineController"] = append(beego.GlobalControllerRouter["web/controllers:MachineController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/machine/index`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:MachineController"] = append(beego.GlobalControllerRouter["web/controllers:MachineController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/machine/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:MachineController"] = append(beego.GlobalControllerRouter["web/controllers:MachineController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/machine/edit`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:MachineController"] = append(beego.GlobalControllerRouter["web/controllers:MachineController"],
		beego.ControllerComments{
			Method: "Remove",
			Router: `/machine/remove`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:MainController"] = append(beego.GlobalControllerRouter["web/controllers:MainController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:MainController"] = append(beego.GlobalControllerRouter["web/controllers:MainController"],
		beego.ControllerComments{
			Method: "Export",
			Router: `/main/export`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:TaskController"] = append(beego.GlobalControllerRouter["web/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/task/index`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:TaskController"] = append(beego.GlobalControllerRouter["web/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/task/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:TaskController"] = append(beego.GlobalControllerRouter["web/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/task/edit`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:TaskController"] = append(beego.GlobalControllerRouter["web/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Remove",
			Router: `/task/remove`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
