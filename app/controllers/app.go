package controllers

import (
	"github.com/rchargel/goblog/app/domain"
	"github.com/revel/revel"
)

type robot string

func (r robot) Apply(req *revel.Request, resp *revel.Response) {
	resp.Out.Header().Set("Content-Type", "text/plain")
	resp.Out.Write([]byte("User-Agent: *\n"))
	resp.Out.Write([]byte("Allow: "))
	resp.Out.Write([]byte(r))
	resp.Out.Write([]byte("\n"))
}

// App is the basic controller for the website.
type App struct {
	*revel.Controller
}

// Index gets the index page for the website.
func (c App) Index() revel.Result {
	tab := "index"
	return c.Render(tab)
}

// Robots get the robots.txt file.
func (c App) Robots() revel.Result {
	return robot("/")
}

// Bookmarks gets the bookmarks.
func (c App) Bookmarks() revel.Result {
	links := domain.GetBookmarks()
	tab := "bookmarks"
	return c.Render(links, tab)
}
