package domain

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// NewSlideShowMap creates the slideshow mapping
func NewSlideShowMap() *SlideShowMap {
	slideShowMap := &SlideShowMap{}
	slideShowMap.init("./public/resources/slideshow.yml")

	return slideShowMap
}

// SlideShow the pointer to the slideshow.
type SlideShow struct {
	Name       string
	Slug       string
	Footer     string
	Theme      string
	Transition string
}

type slideShowList struct {
	Items []SlideShow
}

// SlideShowMap the mapping for the slideshows.
type SlideShowMap struct {
	Items map[string]SlideShow
}

// GetSlideShow gets the slideshow by slug.
func (s *SlideShowMap) GetSlideShow(slug string) *SlideShow {
	item := s.Items[slug]
	return &item
}

// GetContent gets the slideshow content.
func (s *SlideShow) GetContent() string {
	buffer := &bytes.Buffer{}
	file, err := os.Open("./public/resources/slideshow/" + s.Slug + ".html")
	defer file.Close()
	if err == nil {
		buffer.ReadFrom(file)
		return buffer.String()
	}
	return ""
}

func (s *SlideShowMap) init(filepath string) {
	s.Items = make(map[string]SlideShow, 0)
	file, err := os.Open(filepath)
	defer file.Close()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			type itemListType struct {
				Items []SlideShow
			}
			itemList := itemListType{}
			yaml.Unmarshal(data, &itemList)
			fmt.Println(itemList)
			for _, item := range itemList.Items {
				s.Items[item.Slug] = item
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}
