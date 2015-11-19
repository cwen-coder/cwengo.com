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

type MyLabel struct {
	Name string
	Id   string
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
		//AllLabels := make([]string, 10)
		Lables := strings.Split(topic.Lables, " ")
		var MyLables []MyLabel

		if len(Lables) > 0 {
			for i, n := 0, len(Lables); i < n; i++ {
				//AllLabels[i] = models.GetLabel(Lables[i])
				LabelOne := MyLabel{
					Name: models.GetLabel(Lables[i]),
					Id:   Lables[i],
				}
				MyLables = append(MyLables, LabelOne)
			}
			//this.Data["LablesId"] = Lables[0:len(Lables)]
		}
		this.Data["Label"] = MyLables
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
