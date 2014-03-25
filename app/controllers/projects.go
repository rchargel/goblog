package controllers

import (
	"github.com/rchargel/goblog/app/domain"
	"github.com/robfig/revel"
	"strings"
)

var projectList = domain.NewProjectList()

type Projects struct {
	*revel.Controller
}

func (c Projects) List() revel.Result {
	projects := projectList.GetFilteredList()
	return c.Render(projects)
}

func (c Projects) Show() revel.Result {
	path := c.Request.RequestURI
	action := path[strings.LastIndex(path, "/")+1:]
	project := projectList.GetItem(action)
	if project != nil {
		content := project.GetContent()
		meta_description := project.Abstract
		meta_keywords := project.Tags
		return c.Render(project, content, meta_description, meta_keywords)
	} else {
		return c.NotFound("Could not find page " + c.Action)
	}
}
