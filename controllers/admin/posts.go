package admin

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type PostsController struct {
	beego.Controller
}

func (p *PostsController) New() {
	id := p.Ctx.Input.Param(":id")
	p.Data["id"] = id
	p.TplName = "admin/posts/new.tpl"
}
func (p *PostsController) Show() {
	id := p.Ctx.Input.Param(":id")
	p.Data["id"] = id
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
