package admin

import (
	"cwengo.com/models"
	"github.com/astaxie/beego"
)

type LabelController struct {
	beego.Controller
}

func (this *LabelController) Prepare() {
	userid := this.GetSession("userid")
	username := this.GetSession("username")
	if userid == nil || username == nil {
		this.Ctx.Redirect(302, "/admin/login")
		return
	}
}

func (this *LabelController) Get() {
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/Tpl/T.label.tpl"
	this.Data["isLabel"] = true
	this.Data["Username"] = this.GetSession("username")
	var err error
	this.Data["Labels"], err = models.GetAllLabels()
	if err != nil {
		beego.Error(err)
	}
}

func (this *LabelController) Post() {
	NewLabel := this.Input().Get("NewLabel")
	if NewLabel == "" {
		this.Data["LabelNull"] = "表签名不能为空"
		return
	}
	err := models.AddLabel(NewLabel)
	if err != nil {
		beego.Error(err)
	}
	this.Ctx.Redirect(302, "/admin/label")
}

func (this *LabelController) DelLabel() {
	label_id := this.Input().Get("label_id")
	var data map[string]interface{}
	data = make(map[string]interface{})
	if label_id == "" {
		data["status"] = -1
		data["msg"] = "标签ID不能为空"
		this.Ctx.Output.Json(data, true, true)
		return
	}
	err := models.DelLabel(label_id)
	if err != nil {
		beego.Error(err)
		data["status"] = -1
		data["msg"] = "删除失败"
		this.Ctx.Output.Json(data, true, true)
		return
	}
	data["status"] = 1
	data["msg"] = "删除成功"
	this.Ctx.Output.Json(data, true, true)
	return
}

func (this *LabelController) EditLabel() {
	label_id := this.Input().Get("label_id")
	label_name := this.Input().Get("label_name")
	var data map[string]interface{}
	data = make(map[string]interface{})
	if label_id == "" {
		data["status"] = -1
		data["msg"] = "分类ID不能为空"
		this.Ctx.Output.Json(data, true, true)
		return
	}
	if label_name == "" {
		data["status"] = -1
		data["msg"] = "分类名不能为空"
		this.Ctx.Output.Json(data, true, true)
		return
	}
	err := models.EditLabel(label_id, label_name)
	if err != nil {
		data["status"] = -1
		data["msg"] = err.Error()
		this.Ctx.Output.Json(data, true, true)
		return
	}
	data["status"] = 1
	data["msg"] = "修改成功"
	this.Ctx.Output.Json(data, true, true)
	return
}
