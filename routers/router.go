package routers

import (
	"portale/controllers"
    admin "portale/modules/admin/controllers"
    _ "portale/modules/automezzi/controllers"
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
    //beego.Router("/edit/", &controllers.MainController{}, "get,post:Profile")
   
    //*** modules admin ***//
	//beego.Router("/admin/index/:parms", &admin.AdminController{}, "get,post:Index")
	beego.Router("/admin/add/:parms", &admin.AdminController{}, "get,post:Add")
	//beego.Router("/admin/update/:username", &admin.AdminController{}, "get,post:Update")
	beego.Router("/admin/", &admin.AdminController{}, "get,post:Manage")    
	beego.Router("/admin/:parms", &admin.AdminController{}, "get,post:Manage")
	beego.Router("/admin/user/:parms", &admin.AdminController{}, "get,post:UsersManage")
    
     //*** modules automezzi ***//   
    //beego.Router("/automezzi/", &automezzi.AutomezziController{}, "get.post:Gestione")

       
}
