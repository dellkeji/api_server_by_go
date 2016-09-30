package routers

import (
	"bee_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/cc_cache",
		beego.NSRouter("/index", &controllers.IndexController{}),
		beego.NSRouter("/set_property/query_by_set_ids/", &controllers.SetProperty{}),
	)
	beego.AddNamespace(ns)
}
