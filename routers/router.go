package routers

import (
	"cmsByBeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    beego.Router("/menu",&controllers.MenuController{},"Get:Index")
    beego.Router("menu/list",&controllers.MenuController{},"*:List")
    beego.Router("menu/edit",&controllers.MenuController{},"*:Edit")
    beego.Router("menu/editdo",&controllers.MenuController{},"*:EditDo")


    // user
	beego.Router("/user",&controllers.UserController{},"Get:Index")
	beego.Router("/user/add",&controllers.UserController{},"Get:Add")
	beego.Router("/user/adddo",&controllers.UserController{},"*:AddDo")
	beego.Router("/user/edit",&controllers.UserController{},"Get:Edit")
	beego.Router("/user/editdo",&controllers.UserController{},"*:EditDo")
	beego.Router("/user/deletedo",&controllers.UserController{},"Get:DeleteDo")
	beego.Router("/user/list",&controllers.UserController{},"*:List")

	//login
	beego.Router("/login",&controllers.LoginController{},":Index")


    //format
    beego.Router("/format/edit",&controllers.FormateController{},"Get:Edit")
    beego.Router("/format/editdo",&controllers.FormateController{},"Get:EditDo")

	beego.Router("data/?:mid",&controllers.DataController{},"Det:Index")
	beego.Router("data/list/?:mid",&controllers.DataController{},"*:List")
}
