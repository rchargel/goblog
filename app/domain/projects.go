package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

type ProjectList struct {
	items []ProjectItem
}

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

func (c *ProjectList) Index() {
	for i, item := range c.items {
		if item.Active {
			c.items[i] = indexProjectItem(item)
		}
	}
}
