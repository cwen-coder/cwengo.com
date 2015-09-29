package admin

import (
	"cwengo.com/models"
	"fmt"
	"github.com/astaxie/beego"
	"reflect"
)

type AddTopicController struct {
	beego.Controller
}

func (this *AddTopicController) Get() {
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/Tpl/T.addTopic.tpl"
	this.Data["isAddTopic"] = true
	this.Data["Username"] = this.GetSession("username")
	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Labels"], err = models.GetAllLabels()
	if err != nil {
		beego.Error(err)
	}
}

func (this *AddTopicController) Post() {
	title := this.Input().Get("title")
	category := this.Input().Get("category")
	label := this.Input().Get("label")
	this.Ctx.WriteString(title)
	this.Ctx.WriteString(category)
	this.Ctx.WriteString(label)
	fmt.Println(reflect.TypeOf(category))
	fmt.Println(reflect.TypeOf(label))
	return
}
