package domain

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

// NewBlogList creates the initialized bloglist
func NewBlogList() *BlogList {
	blog := &BlogList{}
	blog.init()

	return blog
}

// Date the blog posted date
type Date struct {
	Time time.Time
}

// BlogItem a single blog entry
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

// BlogItemByDate a sorted list of blog items.
type BlogItemByDate []BlogItem

// BlogItemFilterActive a filtered list of blog items.
type BlogItemFilterActive []BlogItem

// BlogList the returned bloglist.
type BlogList struct {
	items []BlogItem
}

// UnmarshalJSON reads the bytes into a JSON formatted string.
func (d *Date) UnmarshalJSON(s []byte) error {
	datetime := string(s)[0:len(s)-1] + "T00:00:00Z\""
	return d.Time.UnmarshalJSON([]byte(datetime))
}

func (l BlogItemByDate) Len() int           { return len(l) }
func (l BlogItemByDate) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l BlogItemByDate) Less(i, j int) bool { return l[i].Date.Time.After(l[j].Date.Time) }

// Filter filters blogitems.
func (l BlogItemFilterActive) Filter() []BlogItem {
	newList := make([]BlogItem, 0, len(l))
	for _, item := range []BlogItem(l) {
		if item.Active {
			newList = append(newList, item)
		}
	}
	return newList
}

// GetContent gets the blog content
func (b *BlogItem) GetContent() string {
	buffer := &bytes.Buffer{}
	file, err := os.Open("./public/resources/blog/" + b.File)
	defer file.Close()
	if err == nil {
		buffer.ReadFrom(file)
		return buffer.String()
	}
	return ""
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
		var blogs []BlogItem
		file, err := os.Open("./public/resources/blogs.yml")
		defer file.Close()
		if err == nil {
			data, err := ioutil.ReadAll(file)
			if err == nil {
				m := make(map[string]interface{})
				yaml.Unmarshal(data, &m)
				if list, ok := m["blogs"].([]interface{}); ok {
					blogs = make([]BlogItem, len(list), len(list))
					for i, item := range list {
						if b, ok := item.(map[interface{}]interface{}); ok {
							dateStr := b["date"].(string)
							t, _ := time.Parse("2006-01-02", dateStr)
							blogs[i] = BlogItem{
								Name:     b["name"].(string),
								Date:     Date{Time: t},
								Abstract: b["abstract"].(string),
								Slug:     b["slug"].(string),
								Tags:     b["tags"].(string),
								Active:   b["active"].(bool),
								File:     fmt.Sprintf("%v.html", b["slug"]),
							}
						}
					}
				}
			}
		} else {
			panic(err)
		}
		sort.Sort(BlogItemByDate(blogs))
		c.items = blogs
	}
}

// Index used by the controller to output the index.
func (c *BlogList) Index() {
	for i, item := range c.items {
		if item.Active {
			c.items[i] = indexBlogItem(item)
		}
	}
}

// GetFilteredList gets a filtered list of blogitems.
func (c *BlogList) GetFilteredList() []BlogItem {
	items := c.items
	return BlogItemFilterActive(items).Filter()
}

// GetItem gets a single blog item.
func (c *BlogList) GetItem(slug string) (*BlogItem, string, string) {
	for i, item := range c.items {
		if item.Slug == slug {
			var prev, next string
			if i > 0 {
				n := i - 1
				for {
					if n == -1 {
						next = ""
						break
					}
					if c.items[n].Active {
						next = c.items[n].Slug
						break
					}
					n--
				}
			}
			if i < (len(c.items) - 1) {
				p := i + 1
				for {
					if p == len(c.items) {
						prev = ""
						break
					}
					if c.items[p].Active {
						prev = c.items[p].Slug
						break
					}
					p++
				}
			}
			return &item, prev, next
		}
	}
	return nil, "", ""
}
