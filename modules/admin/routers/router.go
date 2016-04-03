package routers

import (
	"portale/modules/admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
   	/*beego.Router("/admin/index/:parms", &controllers.AdminController{}, "get,post:Index")
	beego.Router("/admin/add/:parms", &controllers.AdminController{}, "get,post:Add")
	beego.Router("/admin/update/:username", &controllers.AdminController{}, "get,post:Update")
	beego.Router("/admin/", &controllers.AdminController{}, "get,post:Manage")    
	beego.Router("/admin/manage/:parms", &controllers.AdminController{}, "get,post:Manage")
	beego.Router("/admin/manager/user/:parms", &controllers.AdminController{}, "get,post:UsersManage")*/

    
    ns :=
    beego.NewNamespace("/admin",
        beego.NSCond(func(ctx *context.Context) bool {
            if sessionMap, ok := ctx.Input.Session("portale").(map[string]interface{}); !ok {
			beego.Debug("Session: ", sessionMap)
			if sessionMap["admin"] != 3 {
                return true
			}
            return false
		}
        }),
        //beego.NSBefore(auth),
        beego.NSRouter("/index/:parms", &admin.AdminController{}, "get,post:Index")
        beego.NSRouter("/add/:parms", &controllers.AdminController{}, "get,post:Add")
        beego.NSRouter("/update/:username", &controllers.AdminController{}, "get,post:Update")
        beego.NSRouter("/manage/", &controllers.AdminController{}, "get,post:Manage")    
        beego.NSRouter("/manage/:parms", &controllers.AdminController{}, "get,post:Manage")
        beego.NSRouter("/manager/user/:parms", &controllers.AdminController{}, "get,post:UsersManage")
        ),
    )



    //register namespace


}
