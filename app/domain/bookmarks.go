package domain

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Bookmark a bookmark record
type Bookmark struct {
	Link     string
	LinkText string
	Message  string
}

type bookmarks struct {
	Bookmarks []Bookmark
}

var bookmarkList bookmarks

// GetBookmarks gets the bookmarks
func GetBookmarks() []Bookmark {
	if bookmarkList.Bookmarks == nil {
		bookmarkList, _ = readBookmarks("./public/resources/bookmarks.yml")
	}
	return bookmarkList.Bookmarks
}

func readBookmarks(filepath string) (bookmarks, error) {
	bks := bookmarks{}
	file, err := os.Open(filepath)
	defer file.Close()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			yaml.Unmarshal(data, &bks)
		}
	}
	return bks, err
}
