package routers

import (
	"blog/controllers"
	"blog/controllers/admin"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/posts/:id", &controllers.PostsController{}, "get:Show")

	beego.Router("/admin/posts/:id", &admin.PostsController{}, "put:Update")
	beego.Router("/admin/posts/new", &admin.PostsController{}, "get:New")
	beego.Router("/admin/posts", &admin.PostsController{}, "post:Create")
	beego.Router("/admin/posts/:id", &admin.PostsController{}, "delete:Destroy")
}
