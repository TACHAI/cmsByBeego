package main

import (
	_ "cmsByBeego/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.SetLevel(beego.LevelInformational)
	logs.SetLogger("file",`{"filename":"logs/test.log"}`)
	beego.Run()
}

