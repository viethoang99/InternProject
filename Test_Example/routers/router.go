package routers

import (
	"Test_Example/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/department",
			beego.NSInclude(
				&controllers.DepartmentController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
