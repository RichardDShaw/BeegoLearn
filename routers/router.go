package routers

import (
	"myapp/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Get("/get",func(cxt *context.Context){
    	cxt.Output.Body([]byte("Get"))
	})
	beego.Router("/path:id([0-9]+)",&controllers.MainController{},
						"get:Fun1")
 	beego.Router("/path:id:string",&controllers.MainController{},
 						"get:Fun2")
}
