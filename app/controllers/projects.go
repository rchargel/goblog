package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/robfig/revel"
	"io"
	"os"
	"sort"
	"strings"
)

type ProjectItem struct {
	Name     string
	Slug     string
	Abstract string
	File     string
	Active   bool
}

type Projects struct {
	*revel.Controller
}

type ByName []ProjectItem

func (l ByName) Len() int           { return len(l) }
func (l ByName) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByName) Less(i, j int) bool { return l[i].Name < l[j].Name }

func (c *ProjectItem) getContent() string {
	buffer := &bytes.Buffer{}
	file, err := os.Open("./public/resources/projects/" + c.File)
	defer file.Close()
	if err == nil {
		buffer.ReadFrom(file)
		return buffer.String()
	} else {
		return ""
	}
}

func (c Projects) items() []ProjectItem {
	projects := make([]ProjectItem, 0, 15)
	file, err := os.Open("./public/resources/projects.json")
	defer file.Close()
	if err == nil {
		decoder := json.NewDecoder(file)
		for {
			if err := decoder.Decode(&projects); err == nil {
				// resume will get set
			} else if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
	} else {
		panic(err)
	}
	sort.Sort(ByName(projects))
	return projects
}

func (c Projects) item(slug string) *ProjectItem {
	for _, item := range c.items() {
		if item.Slug == slug {
			return &item
		}
	}
	return nil
}

func (c Projects) List() revel.Result {
	projects := c.items()
	return c.Render(projects)
}

func (c Projects) Show() revel.Result {
	path := c.Request.RequestURI
	action := path[strings.LastIndex(path, "/")+1:]
	project := c.item(action)
	if project != nil {
		content := project.getContent()
		return c.Render(project, content)
	} else {
		return c.NotFound("Could not find page " + c.Action)
	}
}
