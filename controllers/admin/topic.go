package admin

import (
	"cwengo.com/models"
	/*"fmt"*/
	/*"reflect"*/
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Prepare() {
	userid := this.GetSession("userid")
	username := this.GetSession("username")
	if userid == nil || username == nil {
		this.Ctx.Redirect(302, "/admin/login")
		return
	}
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

func (this *TopicController) EditTopicShow() {
	tid := this.Input().Get("tid")
	this.Data["Username"] = this.GetSession("username")
	this.Data["toplicList"] = true
	if tid == "" {
		this.Layout = "admin/layout.html"
		this.TplNames = "admin/Tpl/T.error.tpl"
	} else {
		this.Layout = "admin/layout.html"
		this.TplNames = "admin/Tpl/T.editTopic.tpl"
		topic, err := models.GetTopic(tid)
		if err != nil {
			beego.Error(err)
			this.Redirect("/admin/home", 302)
			return
		}
		this.Data["Topic"] = topic
		Categories, err := models.GetAllCategories()
		if err != nil {
			beego.Error(err)
		}
		for _, v := range Categories {
			cate, _ := strconv.Atoi(topic.Category)
			if v.Id == int64(cate) {
				v.IsSelected = true
			} else {
				v.IsSelected = false
			}
		}
		this.Data["Categories"] = Categories
		AllLabels, err := models.GetAllLabels()
		Lables := strings.Split(topic.Lables, " ")
		if len(Lables) > 0 {
			for _, v := range AllLabels {
				v.IsSelected = false
				for _, s := range Lables {
					lab, _ := strconv.Atoi(s)
					if v.Id == int64(lab) {
						v.IsSelected = true
						break
					}
				}
			}
		}
		this.Data["Labels"] = AllLabels
		if err != nil {
			beego.Error(err)
		}
	}
}

func (this *TopicController) EditTopicAct() {
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	category := this.Input().Get("category")
	labels := this.GetStrings("label[]")
	content := this.Input().Get("content")
	err := models.EditTopic(tid, title, category, content, labels)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/admin/home", 302)
}
