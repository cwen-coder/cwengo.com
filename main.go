package main

import (
	"cwengo.com/models"
	_ "cwengo.com/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.SessionOn = true
	beego.Run()
}
