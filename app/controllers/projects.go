package controllers

import (
	"github.com/rchargel/goblog/app/domain"
	"github.com/revel/revel"
	"strings"
)

var projectList = domain.NewProjectList()

type Projects struct {
	*revel.Controller
}

func (c Projects) List() revel.Result {
	projects := projectList.GetFilteredList()
	tab := "projects"
	return c.Render(projects, tab)
}

func (c Projects) Show() revel.Result {
	path := c.Request.RequestURI
	action := path[strings.LastIndex(path, "/")+1:]
	project := projectList.GetItem(action)
	tab := "projects"
	if project != nil {
		content := project.GetContent()
		meta_description := project.Abstract
		meta_keywords := project.Tags
		var moreScripts []string
		if len(project.MoreScripts) != 0 {
			moreScripts = make([]string, len(project.MoreScripts))
			for i, script := range project.MoreScripts {
				moreScripts[i] = "js/" + script
			}
		}
		return c.Render(tab, project, content, meta_description, meta_keywords, moreScripts)
	} else {
		return c.NotFound("Could not find page " + c.Action)
	}
}
