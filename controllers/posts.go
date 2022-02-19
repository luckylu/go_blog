package controllers

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type PostsController struct {
	beego.Controller
}

func (p *PostsController) Show() {
	id := p.Ctx.Input.Param(":id")
	p.Data["id"] = id
	p.TplName = "/posts/show.tpl"
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
