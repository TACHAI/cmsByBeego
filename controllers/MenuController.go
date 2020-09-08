package controllers

import (
	"cmsByBeego/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MenuController struct {
	beego.Controller
}

func(c *MenuController)Index(){
	c.LayoutSections=make(map[string]string)
	c.LayoutSections["footerjs"]="menu/footerjs.html"
	c.setTpl("menu/index.html","")
}

func (c *MenuController)List()  {
	data,total:=models.MenuList()
	type MenuEx string{
		models.MenuModel
		ParentName string
	}
	var menu =make(map[int]string)
	for _,v:=range data{
		menu[v.Mid]=v.Name
	}
	var dataEx []MenuEx
	for _,v:=range data{
		dataEx=append(dataEx,MenuEx{*v,menu[v.Parent]})
	}
	c.listJsonResult(consts.JRCodeSucc,"ok",total,dataEx)
}

func (c *MenuController)Add()  {

}

func (c *MenuController)AddDo()  {

}

func (c *MenuController)Edit()  {
	c.Data["Mid"]=c.GetString("mid")
	c.Data["Parent"]=c.GetString("parent")
	c.Data["Seq"]=c.GetString("seq")
	c.Data["Name"]=c.GetString("name")


	var pMenus []models.MenuModel
	data,_:=models.MenuList()
	for _,v:=range data{
		if 0==v.Parent{
			pMenus = append(pMenus,*v)
		}
	}
	c.Data["PMenus"]=pMenus
	c.LayoutSections=make(map[string]string)
	c.LayoutSections["footerjs"]="menu/footerjs_edit.html"
	c.setTpl("menu/edit.html","common/layout_edit.html")
}

func (c *MenuController)EditDo()  {
	var m models.MenuModel
	if err:= c.ParseForm(&m);err==nil{
		orm.NewOrm().Update(&m)
	}
}

