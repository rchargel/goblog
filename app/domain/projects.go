package domain

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

// ProjectList the list of projects
type ProjectList struct {
	Items []ProjectItem
}

// ProjectItem a single project item
type ProjectItem struct {
	Name        string
	Slug        string
	Abstract    string
	File        string
	Tags        string
	Active      bool
	MoreScripts []string
	Index       map[string]int
}

// ProjectItemByName sorts projects by name.
type ProjectItemByName []ProjectItem

// ProjectItemFilterActive filters out non-active projects.
type ProjectItemFilterActive []ProjectItem

// NewProjectList gets a pointer to the list of projects.
func NewProjectList() *ProjectList {
	projects := &ProjectList{}
	projects.init("./public/resources/projects.yml")
	return projects
}

func (c *ProjectList) init(filepath string) {
	if len(c.Items) == 0 {
		fmt.Println("INITIALIZING PROJECTS")
		file, err := os.Open(filepath)
		defer file.Close()
		if err == nil {
			data, err := ioutil.ReadAll(file)
			if err == nil {
				yaml.Unmarshal(data, c)
			}
		} else {
			panic(err)
		}
		sort.Sort(ProjectItemByName(c.Items))
	}
}

// GetFilteredList gets a filtered list of project items.
func (c *ProjectList) GetFilteredList() []ProjectItem {
	return ProjectItemFilterActive(c.Items).Filter()
}

// GetItem gets a single item by its slug
func (c *ProjectList) GetItem(slug string) *ProjectItem {
	for _, item := range c.Items {
		if item.Slug == slug {
			return &item
		}
	}
	return nil
}

// GetContent gets the content from the project item.
func (c *ProjectItem) GetContent() string {
	buffer := &bytes.Buffer{}
	file, err := os.Open("./public/resources/projects/" + c.File)
	defer file.Close()
	if err == nil {
		buffer.ReadFrom(file)
		return buffer.String()
	}
	return ""
}

func (l ProjectItemByName) Len() int           { return len(l) }
func (l ProjectItemByName) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ProjectItemByName) Less(i, j int) bool { return l[i].Name < l[j].Name }

// Filter filters out the non-active items
func (l ProjectItemFilterActive) Filter() []ProjectItem {
	newList := make([]ProjectItem, 0, len(l))
	for _, item := range []ProjectItem(l) {
		if item.Active {
			newList = append(newList, item)
		}
	}
	return newList
}

func indexProjectItem(b ProjectItem) ProjectItem {
	badChars, _ := regexp.Compile("[^\\w|\\s]")
	htmlChars, _ := regexp.Compile("<[^>]+>")
	content := badChars.ReplaceAllString(htmlChars.ReplaceAllString(b.GetContent(), ""), "")
	keys := badChars.ReplaceAllString(b.Name+" "+b.Tags+" "+b.Abstract, "")
	index := make(map[string]int)
	for _, word := range strings.Fields(content) {
		if len(word) > 3 {
			index[strings.ToLower(word)]++
		}
	}
	for _, word := range strings.Fields(keys) {
		index[strings.ToLower(word)]++
	}
	return ProjectItem{
		Name:     b.Name,
		Slug:     b.Slug,
		Abstract: b.Abstract,
		File:     b.File,
		Tags:     b.Tags,
		Active:   b.Active,
		Index:    index,
	}
}

// Index re-indexes the projects
func (c *ProjectList) Index() {
	for i, item := range c.Items {
		if item.Active {
			c.Items[i] = indexProjectItem(item)
		}
	}
}
