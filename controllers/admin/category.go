package admin

import (
	"cwengo.com/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Prepare() {
	userid := this.GetSession("userid")
	username := this.GetSession("username")
	if userid == nil || username == nil {
		this.Ctx.Redirect(302, "/admin/login")
		return
	}
}

func (this *CategoryController) Get() {
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/Tpl/T.category.tpl"
	this.Data["isCategory"] = true
	this.Data["Username"] = this.GetSession("username")
	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}

func (this *CategoryController) Post() {
	newCategory := this.Input().Get("NewCategory")
	if newCategory == "" {
		this.Data["CategoryNull"] = "分类名不能为空"
		return
	}
	err := models.AddCategory(newCategory)
	if err != nil {
		beego.Error(err)
	}
	this.Ctx.Redirect(302, "/admin/category")
}

func (this *CategoryController) DelCategory() {
	category_id := this.Input().Get("category_id")
	var data map[string]interface{}
	data = make(map[string]interface{})
	if category_id == "" {
		data["status"] = -1
		data["msg"] = "分类ID不能为空"
		this.Ctx.Output.Json(data, true, true)
		return
	}
	err := models.DelCategory(category_id)
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

func (this *CategoryController) EditCategory() {
	category_id := this.Input().Get("category_id")
	category_name := this.Input().Get("category_name")
	var data map[string]interface{}
	data = make(map[string]interface{})
	if category_id == "" {
		data["status"] = -1
		data["msg"] = "分类ID不能为空"
		this.Ctx.Output.Json(data, true, true)
		return
	}
	if category_name == "" {
		data["status"] = -1
		data["msg"] = "分类名不能为空"
		this.Ctx.Output.Json(data, true, true)
		return
	}
	err := models.EditCategory(category_id, category_name)
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
