package controllers

import (
	models "blog/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

func (base *BaseController) Prepare() {
	user_id := base.GetSession("user_id")
	if user_id != nil {
		o := orm.NewOrm()
		user := models.Users{Id: user_id.(int64)}
		err := o.Read(&user)
		if err == nil {
			base.Data["user"] = user
		}
	}
}
