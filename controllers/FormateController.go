package controllers

import (
	"cmsByBeego/models"
	"github.com/astaxie/beego/orm"
)

type FormateController struct {
	BaseController
}

func (c *FormateController)Edit()  {
	midvalue,_:=c.GetInt("mid")
	menu:=models.MenuModel{Mid:midvalue}
	orm.NewOrm().Read(&menu)

	c.Data["Mid"]=midvalue
	c.Data["Format"]=menu.Format
	c.LayoutSections=make(map[string]string)
	c.LayoutSections["footerjs"]="format/footerjs_edit.html"
	c.setTpl("format/edit.html","common/layout_edit.html")
}

func (c *FormateController)EditDo()  {
	mid,_:=c.GetInt("mid")
	f:=c.GetString("formatstr")//提交的format信息

	if 0!=mid{
		menu:=models.MenuModel{Mid:mid,Format:f}
		mid,_:=orm.NewOrm().Update(&menu,"format")
		c.jsonResult(200,"ok",mid)
	}

	c.jsonResult(400,"",0)
}
