package routers

import (
	"cwengo.com/controllers"
	"cwengo.com/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/home", &admin.HomeController{})
	beego.Router("/admin/login", &admin.LoginController{})
	beego.Router("/admin/category", &admin.CategoryController{})
	beego.Router("/admin/category/DelCategory", &admin.CategoryController{}, "post:DelCategory")
	beego.Router("/admin/category/EditCategory", &admin.CategoryController{}, "post:EditCategory")
	beego.Router("/admin/label", &admin.LabelController{})
	beego.Router("/admin/label/DelLabel", &admin.LabelController{}, "post:DelLabel")
	beego.Router("/admin/label/EditLabel", &admin.LabelController{}, "post:EditLabel")
	beego.Router("/admin/topic", &admin.TopicController{})
	beego.Router("/admin/topic/DelTopic", &admin.TopicController{}, "post:DelTopic")
	beego.Router("/admin/topic/EditTopicShow", &admin.TopicController{}, "get:EditTopicShow")
	beego.Router("/admin/topic/EditTopicAct", &admin.TopicController{}, "post:EditTopicAct")
	beego.Router("/topic", &controllers.TopicController{})
}
