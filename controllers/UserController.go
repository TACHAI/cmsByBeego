package controllers

import (
	"cmsByBeego/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"hash/adler32"
)

type UserController struct {
	beego.Controller
}

type JSONS struct {
	Code  int
	Msg   string
	Total int64
	Data  []*models.UserModel

}

func(c *UserController)Index(){
	c.LayoutSections=make(map[string]string)
	c.LayoutSections["footerjs"]="user/footerjs.html"
	c.TplName="user/index.html"
	//c.setTpl("menu/index.html","")
}

func (c *UserController)List()  {
	page,err:=c.GetInt("page")
	if err!=nil{
		page=1
	}
	if page<=0{
		page=1
	}
	size,err:=c.GetInt("limit")
	if err!=nil{
		size=20
	}
	result,count :=models.UserList(size,page)
	data,total:=models.UserList(size,page)



	//c.listJsonResult(consts.JRCodeSucc,"ok",total,dataEx)
	res :=&JSONS{200,"OK",total,data}
	c.Data["json"]=res
	c.ServeJSON()
}

func (c *UserController)Add()  {

}

func (c *UserController)AddDo()  {

}

func (c *UserController)Edit()  {
	userid,_:=c.GetInt("userid")

	o:=orm.NewOrm()
	var user =models.UserModel{UserId:userid}
	o.Read(&user)

	user.Password=""
	c.Data["User"]=user

	authmap :=make(map[int]bool)
	if len(user.AuthStr)>0{
		json.Unmarshal([]byte(user.AuthStr),&authmap)

		for v,_:=range authmap{
			authmap[v]=true
		}

	}


	type Menuitem struct {
		Name string
		Ischecked bool
	}

	menu :=models.ParentMenuList()
	menus :=make(map[int]Menuitem)
	for _,v:=range menu{
		menu[v.Mid]=Menuitem{v.Name:authmap[v.Mid]}
	}
	c.Data["Menus"]=menus

	c.LayoutSections=make(map[string]string)
	c.LayoutSections["footerjs"]="user/footerjs.html"
	c.LayoutSections["layout_edit"]="common/layout_edit.html"
	c.TplName="user/edit.html"
}

func (c *UserController)EditDo()  {

}

func (c *UserController)DeleteDo()  {

}