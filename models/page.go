package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)

type Page struct {
	Id      int
	Website string
	Email   string
}

func init()  {
	orm.RegisterDataBase("default","mysql","root:hzc778209@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterModel(new(Page))
}

func GetPage() Page {
	//rtn := Page{Website:"laishishui.com",Email:"tc1206966083@gmail.com"}
	//return rtn;

	o:=orm.NewOrm()
	p:=Page{Id:1}
	err:=o.Read(&p)
	fmt.Print(err)

	return p
}

func UpdatePage()  {
	p:=Page{Id:1,Email:"myemail_updated"}
	o:=orm.NewOrm()
	o.Update(&p)
}