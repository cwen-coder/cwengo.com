package admin

import (
	"cwengo.com/models"
	/*"fmt"*/
	/*"reflect"*/
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
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

func (this *TopicController) Post() {
	title := this.Input().Get("title")
	category := this.Input().Get("category")
	labels := this.GetStrings("label[]")
	content := this.Input().Get("content")
	err := models.AddTopic(title, category, content, labels)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/admin/topic", 302)
}

func (this *TopicController) DelTopic() {
	topic_id := this.Input().Get("topic_id")
	var data map[string]interface{}
	data = make(map[string]interface{})
	if topic_id == "" {
		data["status"] = -1
		data["msg"] = "博客Id不能为空"
		this.Ctx.Output.Json(data, true, true)
		return
	}
	err := models.DelTopic(topic_id)
	if err != nil {
		beego.Error(err)
		data["status"] = -1
		data["msg"] = "删除出现错误"
		this.Ctx.Output.Json(data, true, true)
		return
	}
	data["status"] = 1
	data["msg"] = "删除成功"
	this.Ctx.Output.Json(data, true, true)
	return
}
