package admin

import (
	"blog/models"
	"fmt"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

type PostsController struct {
	beego.Controller
}

func (p *PostsController) New() {
	p.TplName = "admin/posts/new.tpl"
}

func (p *PostsController) Create() {
	o := orm.NewOrm()

	use1r := models.Users{Name: "Michael"}
	_, _, _ = o.ReadOrCreate(&use1r, "Name")

	title := p.GetString("title")
	content := p.GetString("content")
	user := models.Users{Id: 1}
	post := models.Posts{User: &user, Title: title, Content: content}
	if _, err := o.Insert(&post); err != nil {
		fmt.Println(err)
	}
	p.Redirect("/admin/posts", 302)
}

func (p *PostsController) Index() {
	var posts []*models.Posts
	o := orm.NewOrm()
	_, err := o.QueryTable("posts").OrderBy("-created_at").All(&posts)
	if err != nil {
		//handle err
	}
	p.Data["posts"] = &posts
	p.TplName = "admin/posts/index.tpl"
}
func (p *PostsController) Show() {
	id := p.Ctx.Input.Param(":id")
	p.Data["id"] = id
}

func (p *PostsController) Update() {
	id := p.Ctx.Input.Param(":id")
	title := p.GetString("title")
	content := p.GetString("content")
	intid, _ := strconv.Atoi(id)
	post := models.Posts{Id: intid}
	post.Title = title
	post.Content = content
	o := orm.NewOrm()
	if _, err := o.Update(&post, "Title", "Content"); err != nil {
		fmt.Println(err)
	}
	p.Redirect("/admin/posts", 302)
}

func (p *PostsController) Edit() {
	id := p.Ctx.Input.Param(":id")
	intid, _ := strconv.Atoi(id)
	o := orm.NewOrm()
	post := models.Posts{Id: intid}
	o.Read(&post)
	p.Data["post"] = &post
	p.TplName = "admin/posts/edit.tpl"
}

func (p *PostsController) Destroy() {
	id := p.Ctx.Input.Param(":id")
	intid, _ := strconv.Atoi(id)
	o := orm.NewOrm()
	o.Delete(&models.Posts{Id: intid})
	p.Redirect("/admin/posts", 302)
}
