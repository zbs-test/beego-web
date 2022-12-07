package controllers

import (
	"github.com/astaxie/beego"
)

type HeyController struct {
	beego.Controller
}

func (c *HeyController) Get() {
	c.TplName = "index.html"
}
