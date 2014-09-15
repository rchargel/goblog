package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"io"
	"os"
)

type Robot string

type Bookmark struct {
	Link     string
	LinkText string
	Message  string
}

func (r Robot) Apply(req *revel.Request, resp *revel.Response) {
	resp.Out.Header().Set("Content-Type", "text/plain")
	resp.Out.Write([]byte("User-Agent: *\n"))
	resp.Out.Write([]byte("Allow: "))
	resp.Out.Write([]byte(r))
	resp.Out.Write([]byte("\n"))
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	tab := "index"
	return c.Render(tab)
}

func (c App) Robots() revel.Result {
	return Robot("/")
}

func (c App) Game() revel.Result {
	tab := "index"
	moreScripts := []string{"js/phaser.min.js"}
	return c.Render(tab, moreScripts)
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
	tab := "bookmarks"
	return c.Render(links, tab)
}
