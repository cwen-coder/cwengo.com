package controllers

import (
	"cwengo.com/models"
	"github.com/astaxie/beego"
)

type AboutController struct {
	beego.Controller
}

func (this *AboutController) Get() {
	this.Data["IsAbout"] = true
	this.Data["Categories"], _ = models.GetAllCategories()
	this.Data["Labels"], _ = models.GetAllLabels()
	this.Data["NewTopics"], _ = models.GetAllNewTopics()
	this.Data["VIewsTopics"], _ = models.GetAllViewsTopics()
	this.TplNames = "about.html"
}
