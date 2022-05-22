package controllers

import (
	models "blog/models"
	"net/url"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type AuthenticateController struct {
	beego.Controller
}

func (auth *AuthenticateController) Prepare() {
	// id := auth.Ctx.Input.Param(":id")
	return_to := auth.Ctx.Request.Referer()
	user_id := auth.GetSession("user_id")
	return_url := url.PathEscape(return_to)
	if user_id == nil {
		auth.Redirect("/login?return_to="+return_url, 302)
	} else {
		o := orm.NewOrm()
		user := models.Users{Id: user_id.(int64)}
		err := o.Read(&user)
		if err == nil {
			auth.Data["user"] = user
		}
	}
}
