package routers

import (
	"portale/controllers"
    admin "portale/modules/admin/routers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login/", &controllers.MainController{}, "get,post:Login")
    beego.Router("/login/:back", &controllers.MainController{}, "get,post:Login")
    beego.Router("/register", &controllers.MainController{}, "get,post:Register")
    beego.Router("/check/:uuid", &controllers.MainController{}, "get:Verify")
    beego.Router("/forgot", &controllers.MainController{}, "get,post:Forgot")
    beego.Router("/reset/:uuid", &controllers.MainController{}, "get,post:Reset")
    beego.Router("/notice", &controllers.MainController{}, "get:Notice")
    beego.Router("/logout", &controllers.MainController{}, "get:Logout")
    beego.Router("/profile/", &controllers.MainController{}, "get,post:Profile")
    
    beego.AddNamespace(admin)
}
