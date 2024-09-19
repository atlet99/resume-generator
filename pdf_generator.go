package main

import (
    "github.com/jung-kurt/gofpdf/v2"
    "log"
)

func generateResumePDF(resume Resume, filename string) error {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()

    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, "Resume")

    pdf.Ln(12)

    pdf.SetFont("Arial", "", 12)
    pdf.Cell(40, 10, "Name: " + resume.PersonalDetails.Name)
    pdf.Ln(8)
    pdf.Cell(40, 10, "Email: " + resume.PersonalDetails.Email)
    pdf.Ln(8)
    pdf.Cell(40, 10, "Phone: " + resume.PersonalDetails.Phone)
    pdf.Ln(8)
    pdf.Cell(40, 10, "Location: " + resume.PersonalDetails.Location)

    pdf.Ln(12)
    pdf.SetFont("Arial", "B", 14)
    pdf.Cell(40, 10, "Professional Summary")
    pdf.Ln(8)
    pdf.SetFont("Arial", "", 12)
    pdf.MultiCell(0, 10, resume.ProfessionalSummary)

    pdf.Ln(8)
    pdf.SetFont("Arial", "B", 14)
    pdf.Cell(40, 10, "Key Skills")
    pdf.Ln(8)
    pdf.SetFont("Arial", "", 12)
    for _, skill := range resume.KeySkills.Skills {
        pdf.Cell(40, 10, "- " + skill)
        pdf.Ln(6)
    }

    return pdf.OutputFileAndClose(filename)
}