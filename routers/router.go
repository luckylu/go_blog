package routers

import (
	"blog/controllers"
	"blog/controllers/admin"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/posts/:id", &controllers.PostsController{}, "get:Show")
	beego.Router("/posts/:id/comments", &controllers.CommentsController{}, "post:Create")

	beego.Router("/admin/posts/:id", &admin.PostsController{}, "post:Update")
	beego.Router("/admin/posts/new", &admin.PostsController{}, "get:New")
	beego.Router("/admin/posts", &admin.PostsController{}, "post:Create")
	beego.Router("/admin/posts", &admin.PostsController{}, "get:Index")
	beego.Router("/admin", &admin.PostsController{}, "get:Index")
	beego.Router("/admin/posts/:id/edit", &admin.PostsController{}, "get:Edit")
	beego.Router("/admin/posts/:id/destroy", &admin.PostsController{}, "get:Destroy")
	beego.Router("/admin/comments", &admin.CommentsController{}, "get:Index")
	beego.Router("/admin/comments/:id/destroy", &admin.CommentsController{}, "get:Destroy")

	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/register", &controllers.UsersController{}, "get:ShowRegister;post:HandleRegister")
}
