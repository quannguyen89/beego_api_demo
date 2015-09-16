package routers

import (
	"go_api/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.MainController{})
	beego.Router("/house", &controllers.HouseController{})
}
