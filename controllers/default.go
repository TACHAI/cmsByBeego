package controllers

import (
	"cmsByBeego/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"



	c.SetSession("cmsusername","tachai")
	user := c.GetSession("cmsusername")
	fmt.Print(user)
	logs.Informational("user Jack loged in")

 	m:=models.GetPage()
 	c.Data["Website"]=m.Website
 	c.Data["Email"]=m.Email
	c.TplName = "index.tpl"

}
