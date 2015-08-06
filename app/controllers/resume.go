package controllers

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/rchargel/goblog/app/domain"
	"github.com/revel/revel"
)

// Resume the resume controller.
type Resume struct {
	*revel.Controller
}

// ResumePDF a type used to output the resume into PDF format
type ResumePDF struct {
	resume domain.CV
}

// Apply writes the pdf file.
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

	writeEducation := func(education domain.Education) {
		pdf.SetFont("Arial", "B", 11)
		pdf.Ln(4)
		pdf.CellFormat(0, 15, education.Institution+" - "+education.Year, "", 1, "L", false, 0, "")
		pdf.SetFont("Arial", "", 11)
		pdf.CellFormat(0, 15, education.Location, "", 1, "L", false, 0, "")
		pdf.CellFormat(0, 15, education.Degree, "", 1, "L", false, 0, "")
	}

	writeJob := func(job domain.Job) {
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

// Index the index page for the resume controller.
func (c Resume) Index() revel.Result {
	resume, _ := domain.NewCV("./public/resources/resume.yml")

	tab := "resume"
	return c.Render(resume, tab)
}

// Pdf creates the pdf version of the resume.
func (c Resume) Pdf() revel.Result {
	resume, _ := domain.NewCV("./public/resources/resume.yml")

	return ResumePDF{resume}
}

func cleanStr(text string) string {
	spaceRpl, _ := regexp.Compile("\\s+")
	trimRpl, _ := regexp.Compile("^\\s*|\\s*$")

	return trimRpl.ReplaceAllString(spaceRpl.ReplaceAllString(text, " "), "")
}
