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

}
