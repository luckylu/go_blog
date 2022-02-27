package controllers

import (
	models "blog/models"
	"fmt"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type PostsController struct {
	beego.Controller
}

func (p *PostsController) Show() {
	id := p.Ctx.Input.Param(":id")
	intid, _ := strconv.Atoi(id)
	o := orm.NewOrm()
	post := models.Posts{Id: intid}
	o.Read(&post)
	p.Data["post"] = &post
	var Comments []*models.Comments
	o.QueryTable("comments").Filter("Post", intid).RelatedSel().All(&Comments)
	p.Data["comments"] = &Comments
	p.TplName = "posts/show.tpl"
}

func (p *PostsController) Update() {
	id := p.Ctx.Input.Param(":id")
	fmt.Println(id)
}

func (p *PostsController) Destroy() {
	id := p.Ctx.Input.Param(":id")
	fmt.Println(id)
}

func (p *PostsController) Create() {
	id := p.Ctx.Input.Param(":id")
	fmt.Println(id)
}
