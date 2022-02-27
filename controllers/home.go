package controllers

import (
	models "blog/models"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	var posts []*models.Posts
	o := orm.NewOrm()
	_, err := o.QueryTable("posts").OrderBy("-created_at").All(&posts)
	if err != nil {
		fmt.Println(err)
	}
	c.Data["Posts"] = &posts
	c.TplName = "index.tpl"
}
