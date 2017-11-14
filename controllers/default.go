package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Ctx.WriteString("Hello World")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Username"] = "Richard"
	c.TplName = "Html.html"
}
func (c *MainController) Fun1() {
	c.Ctx.WriteString("fun1")
}
func (c *MainController) Fun2() {
	c.Ctx.WriteString("fun2")
}
