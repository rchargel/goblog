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
	"time"
)

func NewBlogList() *BlogList {
	blog := &BlogList{}
	blog.init()

	return blog
}

type Date struct {
	Time time.Time
}

type BlogItem struct {
	Name     string
	Date     Date
	Slug     string
	Abstract string
	Tags     string
	File     string
	Active   bool
	Index    map[string]int
}

type BlogItemByDate []BlogItem

type BlogItemFilterActive []BlogItem

type BlogList struct {
	items []BlogItem
}

func (d *Date) UnmarshalJSON(s []byte) error {
	datetime := string(s)[0:len(s)-1] + "T00:00:00Z\""
	return d.Time.UnmarshalJSON([]byte(datetime))
}

func (l BlogItemByDate) Len() int           { return len(l) }
func (l BlogItemByDate) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l BlogItemByDate) Less(i, j int) bool { return l[i].Date.Time.After(l[j].Date.Time) }

func (l BlogItemFilterActive) Filter() []BlogItem {
	newList := make([]BlogItem, 0, len(l))
	for _, item := range []BlogItem(l) {
		if item.Active {
			newList = append(newList, item)
		}
	}
	return newList
}

func (b *BlogItem) GetContent() string {
	buffer := &bytes.Buffer{}
	file, err := os.Open("./public/resources/blog/" + b.File)
	defer file.Close()
	if err == nil {
		buffer.ReadFrom(file)
		return buffer.String()
	} else {
		return ""
	}
}

func indexBlogItem(b BlogItem) BlogItem {
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
	return BlogItem{
		Name:     b.Name,
		Date:     b.Date,
		Slug:     b.Slug,
		Abstract: b.Abstract,
		Tags:     b.Tags,
		File:     b.File,
		Active:   b.Active,
		Index:    index,
	}
}

func (c *BlogList) init() {
	if len(c.items) == 0 {
		fmt.Println("INITIALIZE BLOG")
		blogs := make([]BlogItem, 0, 15)
		file, err := os.Open("./public/resources/blogs.json")
		defer file.Close()
		if err == nil {
			decoder := json.NewDecoder(file)
			for {
				if err := decoder.Decode(&blogs); err == nil {
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
		sort.Sort(BlogItemByDate(blogs))
		c.items = blogs
	}
}

func (c *BlogList) Index() {
	for i, item := range c.items {
		if item.Active {
			c.items[i] = indexBlogItem(item)
		}
	}
}

func (c *BlogList) GetFilteredList() []BlogItem {
	items := c.items
	return BlogItemFilterActive(items).Filter()
}

func (c *BlogList) GetItem(slug string) (*BlogItem, string, string) {
	for i, item := range c.items {
		if item.Slug == slug {
			var prev, next string
			if i > 0 {
				next = c.items[i-1].Slug
			}
			if i < (len(c.items) - 1) {
				prev = c.items[i+1].Slug
			}
			return &item, prev, next
		}
	}
	return nil, "", ""
}
