package controllers

import (
	"github.com/rchargel/goblog/app/domain"
	"github.com/robfig/revel"
	"strings"
)

var blogList = domain.NewBlogList()

type Blogs struct {
	*revel.Controller
}

func (c Blogs) List() revel.Result {
	blogs := blogList.GetFilteredList()
	return c.Render(blogs)
}

func (c Blogs) Show() revel.Result {
	path := c.Request.RequestURI
	action := path[strings.LastIndex(path, "/")+1:]
	blog, prev, next := blogList.GetItem(action)
	if blog != nil {
		content := blog.GetContent()
		return c.Render(blog, prev, next, content)
	} else {
		return c.NotFound("Could not find blog page " + c.Action)
	}
}
