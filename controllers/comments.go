package controllers

import (
	models "blog/models"
	"fmt"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
)

type CommentsController struct {
	AuthenticateController
}

func (c *CommentsController) Create() {
	id := c.Ctx.Input.Param(":id")
	intid, _ := strconv.Atoi(id)
	content := c.GetString("content")

	o := orm.NewOrm()
	comment := models.Comments{Post: &models.Posts{Id: intid}, User: &models.Users{Id: 1}, Content: content}
	if _, err := o.Insert(&comment); err != nil {
		fmt.Println(err)
	}
	c.Redirect("/posts/"+id, 302)

}
