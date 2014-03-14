package controllers

import (
	"encoding/json"
	"github.com/robfig/revel"
	"io"
	"os"
)

type Bookmark struct {
	Link     string
	LinkText string
	Message  string
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Bookmarks() revel.Result {
	links := make([]Bookmark, 0, 10)
	file, err := os.Open("./public/resources/bookmarks.json")
	defer file.Close()
	if err == nil {
		decoder := json.NewDecoder(file)
		for {
			if err := decoder.Decode(&links); err == nil {
				// links will be set
			} else if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
	} else {
		panic(err)
	}
	return c.Render(links)
}
