package controllers

import (
	"Test_Example/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
)


type UserController struct {
	beego.Controller
}
var userM models.User

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll(){
	defer u.ServeJSON()
	users:=userM.GetAll1(3)
	u.Data["json"]=users
	return
}

// @Title AddUser
// @Description add a new user or users
// @Param body body models.User true "Body content"
// @Success 200 {object} models.Users
// @router / [post]
func (u *UserController) Post(){
	json.Unmarshal(u.Ctx.Input.RequestBody, &userM)
	err:=userM.AddUser()
	defer u.ServeJSON()
	//users:=userM.GetAll()
	u.Data["json"]=err
	return
}

// @Title UpdateUser
// @Description update a user
// @Param body body models.User true "Body content"
// @Success 200 {object} models.User
// @router / [put]
func (u *UserController) Put() {
	defer u.ServeJSON()
	json.Unmarshal(u.Ctx.Input.RequestBody, &userM)

	err := userM.PutItem()
	if err != nil {
		u.Data["json"] =err.Error()
	} else {
		u.Data["json"] = "success"
	}
}

// @Title DeleteUser
// @Description delete a user
// @Param	userid		path 	string	true "user_id"
// @Success 200 {string} delete success!
// @Failure 403 MaSanPham is empty
// @router /:userid [delete]
func (u *UserController) Delete() {
	id := u.GetString(":userid")
	userM.User_id = id
	err := userM.Delete()
	defer u.ServeJSON()
	if err != nil {
		u.Data["json"] = "Lỗi"
		log.Println("err",err)
		return
	}
	u.Data["json"] = "delete success!"
	return
}


