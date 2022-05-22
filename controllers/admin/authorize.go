package admin

import (
	"blog/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type AuthorizeController struct {
	beego.Controller
}

func (auth *AuthorizeController) Prepare() {
	user_id := auth.GetSession("user_id")
	if user_id == nil {
		auth.Redirect("/login", 302)
	}
	o := orm.NewOrm()

	user := models.Users{Id: user_id.(int64)}
	err := o.Read(&user)
	if err == nil {
		if !user.Admin {
			// auth.FlashWrite("没有管理权限", "true")
			auth.Redirect("/", 302)
		}
	}
}
