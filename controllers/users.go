package controllers

import (
	models "blog/models"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

type UsersController struct {
	beego.Controller
}

func (c *UsersController) ShowRegister() {
	c.TplName = "register.tpl"
}

func (c *UsersController) HandleRegister() {
	password := c.GetString("password")
	name := c.GetString("name")
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println(err)
	} else {
		o := orm.NewOrm()
		user := models.Users{Name: name, PasswordDigest: string(bytes)}
		id, err1 := o.Insert(&user)
		if err1 == nil {
			fmt.Println(id)
		} else {
			fmt.Println(err)
		}

	}
	c.Redirect("/login", 302)
}
