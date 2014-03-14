package controllers

import (
	"encoding/xml"
	"github.com/robfig/revel"
	"io"
	"os"
)

type Experience struct {
	Jobs []Job `xml:"Job"`
}

type Education struct {
	Year        string
	Institution string
	Location    string
	Degree      string
}

type Job struct {
	Start       string
	End         string
	Company     string
	Title       string
	Location    string
	Description string
	Skills      string
}

type CV struct {
	Summary             string
	Experience          Experience
	Education           Education
	AdditionalSkills    string
	AdditionalLanguages string
}

type Resume struct {
	*revel.Controller
}

func (c Resume) Index() revel.Result {
	resume := CV{}
	file, err := os.Open("./public/resources/resume.xml")
	defer file.Close()
	if err == nil {
		decoder := xml.NewDecoder(file)
		for {
			if err := decoder.Decode(&resume); err == nil {
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
	return c.Render(resume)
}
