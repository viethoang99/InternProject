package controllers

import (
	"Test_Example/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
)


type DepartmentController struct {
	beego.Controller
}
var DepartmentM models.Department


// @Title GetAll
// @Description get all Departments
// @Success 200 {object} models.Department
// @router / [get]
func (u *DepartmentController) GetAll(){
	defer u.ServeJSON()
	Departments:=DepartmentM.GetAll()
	u.Data["json"]=Departments
	return
}

// @Title AddDepartment
// @Description add a new Department
// @Param body body models.Department true "Body content"
// @Success 200 {object} models.Departments
// @router / [post]
func (u *DepartmentController) Post(){
	json.Unmarshal(u.Ctx.Input.RequestBody, &DepartmentM)
	err:=DepartmentM.AddDepartment()
	defer u.ServeJSON()
	if err != nil {
		u.Data["json"] = err
		return
	}
	Departments:=DepartmentM.GetAll()
	u.Data["json"]=Departments
	return
}

// @Title UpdateDepartment
// @Description update a Department
// @Param body body models.Department true "Body content"
// @Success 200 {object} models.Department
// @router / [put]
func (u *DepartmentController) Put(){
	defer u.ServeJSON()
	json.Unmarshal(u.Ctx.Input.RequestBody, &DepartmentM)
	err := DepartmentM.PutItem()
	if err != nil {
		u.Data["json"] =err.Error()
	} else {
		u.Data["json"] = "success"
	}
}

// @Title DeleteDepartment
// @Description delete a Department
// @Param	Departmentid		path 	string	true "Department_id"
// @Success 200 {string} delete success!
// @Failure 403 MaPhong Ban is empty
// @router / [delete]
func (u *DepartmentController) Delete() {
	id := u.GetString("Departmentid")
	DepartmentM.Department_id=id
	err := DepartmentM.Delete()
	defer u.ServeJSON()
	if err != nil {
		u.Data["json"] = "Lỗi"
		log.Println("err",err)
		return
	}
	u.Data["json"] = "delete success!"
	return
}

// @Title GetUser
// @Description Get user from department
// @Param	Departmentid		path 	string	true "Department_id"
// @Success 200 {object} models.User
// @Failure 403 MaPhongBan is empty
// @router /:Departmentid [get]
func (u *DepartmentController) GetUsers() {
	defer u.ServeJSON()
	mapb := u.GetString(":Departmentid")
	DepartmentM.Department_id=mapb
	if mapb != "" {
		user := userM.GetAll()
		var users []models.User
		for i:=range user{
			if user[i].Department_id==mapb{
				users=append(users,user[i])
			}
		}
		u.Data["json"]=users
	}else{
		u.Data["json"] = "INTERNAL SERVER ERROR"
	}
	return
}

// @Title GetUserSame
// @Description Get user from same department
// @Param	userid		path 	string	true "user_id"
// @Success 200 {object} models.User
// @Failure 403 MaNhanVien is empty
// @router /:userid [get]
func (u *DepartmentController) GetUsersSame() {
	defer u.ServeJSON()
	user_id := u.GetString(":userid")
	userM.User_id=user_id
	if user_id != "" {
		user := userM.GetAll()
		var users []models.User
		for i:=range user{
			mapb:=user[i].Department_id
			if user[i].Department_id==mapb{
				users=append(users,user[i])
			}
		}
		u.Data["json"]=users
	}else{
		u.Data["json"] = "INTERNAL SERVER ERROR"
	}
	return
}

// @Title DeleteUserFromDepartment
// @Description xoa nguoi dung cua mot phong ban
// @Param   userid    query   string false       "user_id"
// @Param   departmentid    query   string false       "department_id"
// @Success 200 string xoa thanh cong
// @Failure 403 INTERNAL SERVER ERROR
// @router /updateuserdepartment [put]
func (u *DepartmentController) UpdateUserFromDepartment() {
	defer u.ServeJSON()
	var user_id,newdep_id string
	u.Ctx.Input.Bind(&user_id, "userid")
	u.Ctx.Input.Bind(&newdep_id, "departmentid")
	userM.User_id = user_id
	userM.Department_id = newdep_id
	string := DepartmentM.UpdateUser(userM)
	u.Data["json"] = string
}


