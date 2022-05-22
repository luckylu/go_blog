package controllers

import (
	models "blog/models"
	"fmt"
	"net/url"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) ShowLogin() {
	c.TplName = "login.tpl"
}

func (c *LoginController) HandleLogin() {
	name := c.GetString("name")
	password := c.GetString("password")
	o := orm.NewOrm()
	user := models.Users{Name: name}
	err := o.Read(&user, "Name")
	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
		fmt.Println(err)
		if err == nil {
			c.SetSession("user_id", user.Id)
			return_to := c.GetString("return_to")
			return_url, _ := url.PathUnescape(return_to)
			if return_to != "" {
				c.Redirect(return_url, 302)
			} else {
				c.Redirect("/", 302)
			}
		} else {
			c.Redirect("/login", 302)
		}
	} else {
		c.Redirect("/login", 302)
	}
}

func (c *LoginController) Logout() {
	c.DelSession("user_id")
	c.Redirect("/", 302)
}
