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
}
