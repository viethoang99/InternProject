package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Test_Example/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["Test_Example/controllers:DepartmentController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Test_Example/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["Test_Example/controllers:DepartmentController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Test_Example/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["Test_Example/controllers:DepartmentController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Test_Example/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["Test_Example/controllers:DepartmentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:Departmentid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Test_Example/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["Test_Example/controllers:DepartmentController"],
        beego.ControllerComments{
            Method: "GetUsers",
            Router: "/:Departmentid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Test_Example/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["Test_Example/controllers:DepartmentController"],
        beego.ControllerComments{
            Method: "GetUsersSame",
            Router: "/:userid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Test_Example/controllers:UserController"] = append(beego.GlobalControllerRouter["Test_Example/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Test_Example/controllers:UserController"] = append(beego.GlobalControllerRouter["Test_Example/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Test_Example/controllers:UserController"] = append(beego.GlobalControllerRouter["Test_Example/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Test_Example/controllers:UserController"] = append(beego.GlobalControllerRouter["Test_Example/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:userid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
