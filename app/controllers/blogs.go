package controllers

import (
	"encoding/json"
	"github.com/robfig/revel"
	"io"
	"os"
	"sort"
	"time"
)

type Blogs struct {
	*revel.Controller
}

type Date struct {
	Time time.Time
}

type BlogItem struct {
	Name     string
	Date     Date
	Slug     string
	Abstract string
	File     string
	Active   bool
}

type ByBlogDate []BlogItem

type FilterBlogActive []BlogItem

func (d *Date) UnmarshalJSON(s []byte) error {
	datetime := string(s)[0:len(s)-1] + "T00:00:00Z\""
	return d.Time.UnmarshalJSON([]byte(datetime))
}

func (l ByBlogDate) Len() int           { return len(l) }
func (l ByBlogDate) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByBlogDate) Less(i, j int) bool { return l[i].Date.Time.After(l[j].Date.Time) }

func (l FilterBlogActive) Filter() []BlogItem {
	newList := make([]BlogItem, 0, len(l))
	for _, item := range []BlogItem(l) {
		if item.Active {
			newList = append(newList, item)
		}
	}
	return newList
}

func (c Blogs) items() []BlogItem {
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
	blogs = FilterBlogActive(blogs).Filter()
	sort.Sort(ByBlogDate(blogs))
	return blogs
}

func (c Blogs) item(slug string) (*BlogItem, string, string) {
	items := c.items()
	for i, item := range items {
		if item.Slug == slug {
			var prev, next string
			if i > 0 {
				prev = items[i-1].Slug
			}
			if i < (len(items) - 1) {
				next = items[i+1].Slug
			}
			return &item, prev, next
		}
	}
	return nil, "", ""
}

func (c Blogs) List() revel.Result {
	blogs := c.items()
	return c.Render(blogs)
}
