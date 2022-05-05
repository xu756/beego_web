package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"myweb/controllers"
)

func init() {
	beego.Router("/", &controllers.Index{})
	beego.Router("/login", &controllers.Login{})
}
