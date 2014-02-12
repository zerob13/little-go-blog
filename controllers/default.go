package controllers

import (
	"fmt"
	"little-go-blog/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["blogs"] = models.GetAll()
	this.Layout = "layout.tpl"
	this.TplNames = "index.tpl"
}

type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {
	//	inputs := this.Input()
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	this.Data["Post"] = models.GetBlog(id)
	fmt.Println(models.GetBlog(id))
	this.Layout = "layout.tpl"
	this.TplNames = "view.tpl"
}

type NewController struct {
	beego.Controller
}

func (this *NewController) Get() {
	this.Layout = "layout.tpl"
	this.TplNames = "new.tpl"
}

func (this *NewController) Post() {
	inputs := this.Input()
	var blog models.Blog
	blog.Title = inputs.Get("title")
	blog.Content = inputs.Get("content")
	blog.Created = time.Now().Format("Jan 2, 2006 at 3:04pm (MST)")
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}

type EditController struct {
	beego.Controller
}

func (this *EditController) Get() {
	//inputs := this.Input()
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.tpl"
	this.TplNames = "new.tpl"
}

func (this *EditController) Post() {
	inputs := this.Input()
	var blog models.Blog
	blog.Id, _ = strconv.Atoi(inputs.Get("id"))
	blog.Title = inputs.Get("title")
	blog.Content = inputs.Get("content")
	blog.Created = time.Now().Format("Jan 2, 2006 at 3:04pm (MST)")
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}

type DeleteController struct {
	beego.Controller
}

func (this *DeleteController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	this.Data["Post"] = models.GetBlog(id)
	this.Ctx.Redirect(302, "/")
}
