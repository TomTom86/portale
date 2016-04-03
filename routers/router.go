package routers

import (
	"portale/controllers"
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
}
