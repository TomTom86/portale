package routers

import (
	"portale/controllers"
	admin "portale/modules/admin/controllers"
	automezzi "portale/modules/automezzi/controllers"

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
	//beego.Router("/automezzi/:parms", &automezzi.AutomezziController{}, "get.post:Gestione")
	beego.Router("/automezzi/view/fornitori/:parms", &automezzi.AutomezziController{}, "get,post:ViewFornitori")
	beego.Router("/automezzi/view/fornitori/", &automezzi.AutomezziController{}, "get,post:ViewFornitori")

	beego.Router("/automezzi/view/contrattiacq/:parms", &automezzi.AutomezziController{}, "get,post:ViewContrattiAcq")
	beego.Router("/automezzi/edit/contrattoacq/:parms", &automezzi.AutomezziController{}, "get,post:EditAcquisto")
	beego.Router("/automezzi/add/acquisto/:parms", &automezzi.AutomezziController{}, "get,post:AddAcquisto")
	//beego.Router("/automezzi/add/contratto/:parms", &automezzi.AutomezziController{}, "get,post:AddContratto")

	beego.Router("/automezzi/view/contrattilea/:parms", &automezzi.AutomezziController{}, "get,post:ViewContrattiLea")
	beego.Router("/automezzi/edit/contrattilea/:parms", &automezzi.AutomezziController{}, "get,post:EditLeasing")
	beego.Router("/automezzi/add/leasing/:parms", &automezzi.AutomezziController{}, "get,post:AddLeasing")

	beego.Router("/automezzi/view/contrattinol/:parms", &automezzi.AutomezziController{}, "get,post:ViewContrattiNol")
	beego.Router("/automezzi/edit/contrattonol/:parms", &automezzi.AutomezziController{}, "get,post:EditNoleggio")
	beego.Router("/automezzi/add/noleggio/:parms", &automezzi.AutomezziController{}, "get,post:AddNoleggio")

	beego.Router("/automezzi/add/automezzo/", &automezzi.AutomezziController{}, "get,post:AddAutomezzo")
	beego.Router("/automezzi/add/automezzo2/", &automezzi.AutomezziController{}, "get,post:AddAutomezzo2")
	beego.Router("/automezzi/add/conducente/", &automezzi.AutomezziController{}, "get,post:AddConducente")
	beego.Router("/automezzi/add/fornitore/", &automezzi.AutomezziController{}, "get,post:AddFornitore")
	beego.Router("/automezzi/add/data", &automezzi.AutomezziController{}, "get,post:AddData")
}
