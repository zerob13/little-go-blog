package controllers

import (
	"fmt"
	"little-go-blog/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	te := models.GetAll()
	var res []models.Blog
	for _, post := range te {
		temp := []byte(post.Content)
		temp2 := blackfriday.MarkdownBasic(temp)
		post.Content = string(temp2)
		res = append(res, post)
	}
	this.Data["blogs"] = res

	this.Layout = "layout.tpl"
	this.TplNames = "index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["MenuTpl"] = "menu/home.tpl"
}

type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {
	//	inputs := this.Input()
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	post := models.GetBlog(id)
	input := []byte(post.Content)
	output := blackfriday.MarkdownBasic(input)
	post.Content = string(output)
	this.Data["Post"] = post
	this.Layout = "layout.tpl"
	this.TplNames = "view.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["MenuTpl"] = "menu/home.tpl"
}

type NewController struct {
	beego.Controller
}

func (this *NewController) Get() {
	this.Layout = "layout.tpl"
	this.TplNames = "new.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["MenuTpl"] = "menu/new.tpl"
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
	this.TplNames = "edit.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["MenuTpl"] = "menu/home.tpl"
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
	//this.Data["Post"] =
	fmt.Println(id)
	models.DelBlog(id)
	this.Ctx.Redirect(302, "/")
}

func CToGoString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}
