package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
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
