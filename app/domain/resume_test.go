package domain

import (
	"strings"
	"testing"
)

/* XML file no longer exists
func Test_NewCV_Xml(t *testing.T) {
	path := "../../public/resources/resume.xml"
	cv, err := NewCV(path)
	if err == nil {
		if cv.validateCV(t) {
			t.Log("Read XML test passed")
		}
	} else {
		t.Errorf("Failed to read file: %s\n", err.Error())
	}
}
*/

func Test_NewCV_Yml(t *testing.T) {
	path := "../../public/resources/resume.yml"
	cv, err := NewCV(path)
	if err == nil {
		if cv.validateCV(t) {
			t.Log("Read YML test passed")
		}
	} else {
		t.Errorf("Failed to read file: %s\n", err.Error())
	}
}

func (cv CV) validateCV(t *testing.T) bool {
	success := true
	if !strings.Contains(cv.Summary, "Senior Software Developer with more than 10 years of experience.") {
		t.Errorf("Does not contain correct summary: %s\n", cv.Summary)
		success = false
	}
	if len(cv.Experience.Jobs) < 6 {
		t.Errorf("Should contain at least 6 jobs but has %v\n", len(cv.Experience.Jobs))
		success = false
	} else {
		job := cv.Experience.Jobs[len(cv.Experience.Jobs)-1]
		if job.Start != "1998" {
			t.Errorf("First job should have started in 1998: %v\n", job.Start)
			success = false
		}
	}
	if cv.Education.Year != "2001" {
		t.Errorf("Should have graduated in 2001: %v\n", cv.Education.Year)
		success = false
	}
	if cv.Education.Institution != "George Mason University" {
		t.Errorf("Invalid institution: %v\n", cv.Education.Institution)
		success = false
	}
	if !strings.Contains(cv.AdditionalLanguages, "Portuguese") {
		t.Errorf("Missing language: %v\n", cv.AdditionalLanguages)
		success = false
	}
	if !strings.Contains(cv.AdditionalSkills, "Grails") {
		t.Errorf("Missing skill: %v\n", cv.AdditionalSkills)
		success = false
	}
	return success
}
