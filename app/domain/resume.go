package domain

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// CV is the curriculum vitae
type CV struct {
	Summary             string
	Experience          Experience
	Education           Education
	AdditionalSkills    string
	AdditionalLanguages string
}

// Experience is a list of jobs
type Experience struct {
	Jobs []Job `xml:"Job"`
}

// Job is the meta-data for one position at a company
type Job struct {
	Start       string
	End         string
	Company     string
	Title       string
	Location    string
	Description string
	Skills      string
}

// Education is the meta-data for education
type Education struct {
	Year        string
	Institution string
	Location    string
	Degree      string
}

// NewCV creates a new CV object
func NewCV(filepath string) (CV, error) {
	cv := CV{}
	var err error
	if strings.Contains(filepath, ".xml") {
		err = initFromXML(&cv, filepath)
	} else if strings.Contains(filepath, ".yml") {
		err = initFromYML(&cv, filepath)
	}
	return cv, err
}

func initFromYML(cv *CV, filepath string) error {
	file, err := os.Open(filepath)
	defer file.Close()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			yaml.Unmarshal(data, cv)
		}
	}
	return err
}

func initFromXML(cv *CV, filepath string) error {
	file, err := os.Open(filepath)
	defer file.Close()
	if err == nil {
		decoder := xml.NewDecoder(file)
		for {
			if err := decoder.Decode(&cv); err == nil {
				// resume will get set
			} else if err == io.EOF {
				break
			} else {
				return err
			}
		}
	}
	return err
}
