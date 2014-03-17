package controllers

import (
	"code.google.com/p/gofpdf"
	"encoding/xml"
	"github.com/robfig/revel"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
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

type ResumePDF struct {
	resume CV
}

func (r ResumePDF) Apply(req *revel.Request, resp *revel.Response) {
	resp.Out.Header().Set("Content-Disposition", `inline; filename="rafael_chargels_resume.pdf"`)

	pdf := gofpdf.New("P", "pt", "letter", "./fonts/")

	header := func(text string) {
		pdf.Ln(4)
		pdf.SetFont("Arial", "B", 16)
		pdf.CellFormat(0, 30, strings.ToUpper(text), "", 1, "L", false, 0, "")
	}

	writeBullet := func(text string) {
		pdf.SetFont("Arial", "", 11)
		pdf.Ln(4)
		pdf.SetFillColor(0, 0, 0)
		pdf.CellFormat(20, 15, "", "", 0, "L", false, 0, "")

		x := pdf.GetX()
		y := pdf.GetY()
		pdf.Circle(x-5, y+7, 2, "F")
		pdf.MultiCell(0, 15, cleanStr(text), "", "L", false)
	}

	writeEducation := func(education Education) {
		pdf.SetFont("Arial", "B", 11)
		pdf.Ln(4)
		pdf.CellFormat(0, 15, education.Institution+" - "+education.Year, "", 1, "L", false, 0, "")
		pdf.SetFont("Arial", "", 11)
		pdf.CellFormat(0, 15, education.Location, "", 1, "L", false, 0, "")
		pdf.CellFormat(0, 15, education.Degree, "", 1, "L", false, 0, "")
	}

	writeJob := func(job Job) {
		pdf.SetFont("Arial", "B", 11)
		pdf.Ln(4)
		pdf.CellFormat(0, 15, job.Start+" - "+job.End, "", 1, "L", false, 0, "")
		pdf.CellFormat(0, 15, job.Title, "", 1, "L", false, 0, "")
		pdf.CellFormat(0, 15, job.Company, "", 1, "L", false, 0, "")
		pdf.SetFont("Arial", "BI", 11)
		pdf.CellFormat(0, 15, job.Location, "", 1, "L", false, 0, "")
		pdf.SetFont("Arial", "", 11)
		pdf.Ln(8)
		pdf.CellFormat(20, 15, "", "", 0, "L", false, 0, "")
		pdf.MultiCell(0, 15, cleanStr(job.Description), "", "L", false)
		pdf.Ln(8)
		pdf.SetFont("Arial", "I", 11)
		pdf.CellFormat(20, 15, "", "", 0, "L", false, 0, "")
		pdf.MultiCell(0, 15, cleanStr(job.Skills), "", "L", false)
		pdf.Ln(8)
	}

	pdf.SetTitle("Rafael Pacheco Chargel - Resume", true)
	pdf.SetAuthor("Rafael Pacheco Chargel", true)
	pdf.SetMargins(30, 30, 30)
	pdf.SetAutoPageBreak(true, 30)
	pdf.AddPage()
	pdf.SetFont("Helvetica", "B", 11)
	pdf.CellFormat(0, 15, "Rafael Pacheco Chargel", "", 1, "R", false, 0, "")
	pdf.SetFont("Helvetica", "", 11)
	pdf.CellFormat(0, 15, "http://zcarioca.net", "", 1, "R", false, 0, "")
	header("Summary")

	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(0, 15, cleanStr(r.resume.Summary), "", "L", false)

	header("Experience")
	for _, job := range r.resume.Experience.Jobs {
		writeJob(job)
	}

	header("Other Skills")
	writeBullet("Skills: " + r.resume.AdditionalSkills)
	writeBullet("Languages: " + r.resume.AdditionalLanguages)

	header("Education")
	writeEducation(r.resume.Education)

	resp.WriteHeader(http.StatusOK, "application/pdf")
	pdf.Output(resp.Out)
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

func (c Resume) Pdf() revel.Result {
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
	return ResumePDF{resume}
}

func cleanStr(text string) string {
	spaceRpl, _ := regexp.Compile("\\s+")
	trimRpl, _ := regexp.Compile("^\\s*|\\s*$")

	return trimRpl.ReplaceAllString(spaceRpl.ReplaceAllString(text, " "), "")
}
