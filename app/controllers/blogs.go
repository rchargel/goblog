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
	tab := "blog"
	return c.Render(blogs, tab)
}

func (c Blogs) Show() revel.Result {
	const dateFormat = "Jan. 1, 2006"
	path := c.Request.RequestURI
	action := path[strings.LastIndex(path, "/")+1:]
	blog, prev, next := blogList.GetItem(action)
	tab := "blog"
	if blog != nil {
		content := blog.GetContent()
		meta_description := blog.Abstract
		meta_date := blog.Date.Time.Format(dateFormat)
		meta_keywords := blog.Tags
		return c.Render(tab, blog, prev, next, content, meta_description, meta_date, meta_keywords)
	} else {
		return c.NotFound("Could not find blog page " + c.Action)
	}
}
