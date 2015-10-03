package controllers

import (
	"cwengo.com/models"
	"github.com/astaxie/beego"
	/*"strconv"*/
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	topicId := this.Input().Get("topicId")
	if topicId == "" {
		this.TplNames = "error.html"
	} else {
		topic, err := models.GetTopic(topicId)
		if err != nil {
			beego.Error(err)
			this.Redirect("/", 302)
			return
		}
		this.Data["Topic"] = topic
		this.Data["Category"] = models.GetCategory(topic.Category)
		AllLabels := make([]string, 10)
		Lables := strings.Split(topic.Lables, " ")
		if len(Lables) > 0 {
			for i, n := 0, len(Lables); i < n-1; i++ {
				AllLabels[i] = models.GetLabel(Lables[i])
			}
			this.Data["LablesId"] = Lables[0 : len(Lables)-1]
		}
		this.Data["Label"] = AllLabels
		if err != nil {
			beego.Error(err)
		}
		this.Data["Categories"], _ = models.GetAllCategories()
		this.Data["Labels"], err = models.GetAllLabels()
		this.Data["NewTopics"], err = models.GetAllNewTopics()
		this.Data["VIewsTopics"], err = models.GetAllViewsTopics()
		this.TplNames = "topicView.html"
	}
}
