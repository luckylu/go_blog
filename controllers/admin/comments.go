package admin

import (
	"blog/models"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
)

type CommentsController struct {
	AuthorizeController
}

func (p *CommentsController) Index() {
	var comments []*models.Comments
	o := orm.NewOrm()
	_, err := o.QueryTable("comments").OrderBy("-created_at").All(&comments)
	if err != nil {
		//handle err
	}
	p.Data["comments"] = &comments
	p.TplName = "admin/comments/index.tpl"
}

func (p *CommentsController) Destroy() {
	id := p.Ctx.Input.Param(":id")
	intid, _ := strconv.Atoi(id)
	o := orm.NewOrm()
	o.Delete(&models.Comments{Id: intid})
	p.Redirect("/admin/comments", 302)
}
