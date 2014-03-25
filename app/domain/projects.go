package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type ProjectList struct {
	items []ProjectItem
}

type ProjectItem struct {
	Name     string
	Slug     string
	Abstract string
	File     string
	Tags     string
	Active   bool
}

type ProjectItemByName []ProjectItem

type ProjectItemFilterActive []ProjectItem

func NewProjectList() *ProjectList {
	projects := &ProjectList{}
	projects.Init()
	return projects
}

func (c *ProjectList) Init() {
	if len(c.items) == 0 {
		fmt.Println("INITIALIZING PROJECTS")
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
		sort.Sort(ProjectItemByName(projects))
		c.items = projects
	}
}

func (c *ProjectList) GetFilteredList() []ProjectItem {
	return ProjectItemFilterActive(c.items).Filter()
}

func (c *ProjectList) GetItem(slug string) *ProjectItem {
	for _, item := range c.items {
		if item.Slug == slug {
			return &item
		}
	}
	return nil
}

func (c *ProjectItem) GetContent() string {
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

func (l ProjectItemByName) Len() int           { return len(l) }
func (l ProjectItemByName) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ProjectItemByName) Less(i, j int) bool { return l[i].Name < l[j].Name }

func (l ProjectItemFilterActive) Filter() []ProjectItem {
	newList := make([]ProjectItem, 0, len(l))
	for _, item := range []ProjectItem(l) {
		if item.Active {
			newList = append(newList, item)
		}
	}
	return newList
}
