package controllers

import (
	"github.com/rchargel/goblog/app/domain"
	"github.com/revel/revel"
	"strings"
)

var slideShowMap = domain.NewSlideShowMap()

type SlideShow struct {
	*revel.Controller
}

func (c SlideShow) Show() revel.Result {
	path := c.Request.RequestURI
	action := path[strings.LastIndex(path, "/")+1:]
	if strings.Index(action, "?") != -1 {
		action = action[0:strings.Index(action, "?")]
	}
	slideShow := slideShowMap.GetSlideShow(action)

	if slideShow == nil {
		return c.NotFound("Could not find slide show")
	} else {
		content := slideShow.GetContent()
		return c.Render(slideShow, content)
	}
}
