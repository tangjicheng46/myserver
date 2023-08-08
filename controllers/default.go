package controllers

import (
	"myserver/models"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "myserver"
	c.Data["Email"] = "tangjch15@gmail.com"

	user := models.GetUserById(1)
	if user == nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		c.Ctx.WriteString("Not find user id")
	}
	c.Data["user"] = user

	c.TplName = "index.tpl"
}
