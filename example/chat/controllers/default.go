// Beego (http://beego.me/)
// @description beego is an open-source, high-performance web framework for the Go programming language.
// @link        http://github.com/mvpmvh/beego for the canonical source repository
// @license     http://github.com/mvpmvh/beego/blob/master/LICENSE
// @authors     Unknwon

package controllers

import (
	"github.com/mvpmvh/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["host"] = this.Ctx.Request.Host
	this.TplNames = "index.tpl"
}
