package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func NewSlideShowMap() *SlideShowMap {
	slideShowMap := &SlideShowMap{}
	slideShowMap.init()

	return slideShowMap
}

type SlideShow struct {
	Name       string
	Slug       string
	File       string
	Footer     string
	Theme      string
	Transition string
}

type SlideShowMap struct {
	items map[string]SlideShow
}

func (s *SlideShowMap) GetSlideShow(slug string) *SlideShow {
	item := s.items[slug]
	return &item
}

func (s *SlideShow) GetContent() string {
	buffer := &bytes.Buffer{}
	file, err := os.Open("./public/resources/slideshow/" + s.File)
	defer file.Close()
	if err == nil {
		buffer.ReadFrom(file)
		return buffer.String()
	}
	return ""
}

func (s *SlideShowMap) init() {
	if len(s.items) == 0 {
		fmt.Println("INITIALIZING SLIDE SHOWS")
		slideShows := make([]SlideShow, 0, 15)
		file, err := os.Open("./public/resources/slideshow.json")
		defer file.Close()
		if err == nil {
			decoder := json.NewDecoder(file)
			for {
				if err := decoder.Decode(&slideShows); err == nil {
					// ignore
				} else if err == io.EOF {
					break
				} else {
					panic(err)
				}
			}
		} else {
			panic(err)
		}
		s.items = make(map[string]SlideShow)
		for _, item := range slideShows {
			s.items[item.Slug] = item
		}
	}
}
