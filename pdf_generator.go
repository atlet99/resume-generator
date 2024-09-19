package main

import (
    "github.com/jung-kurt/gofpdf/v2"
    "github.com/joho/godotenv"
    "log"
    "os"
    "strconv"
)

func loadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}

func generateResumePDF(resume Resume, filename string) error {
    loadEnv()

    fontFamily := os.Getenv("FONT_FAMILY")
    fontSizeStr := os.Getenv("FONT_SIZE")
    titleFontSizeStr := os.Getenv("PDF_TITLE_FONT_SIZE")

    fontSize, _ := strconv.ParseFloat(fontSizeStr, 64)
    titleFontSize, _ := strconv.ParseFloat(titleFontSizeStr, 64)

    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()

    pdf.SetFont(fontFamily, "B", titleFontSize)
    pdf.Cell(40, 10, "Resume")

    pdf.Ln(12)

    pdf.SetFont(fontFamily, "", fontSize)
    pdf.Cell(40, 10, "Name: "+resume.PersonalDetails.Name)
    pdf.Ln(8)
    pdf.Cell(40, 10, "Email: "+resume.PersonalDetails.Email)
    pdf.Ln(8)
    pdf.Cell(40, 10, "Phone: "+resume.PersonalDetails.Phone)
    pdf.Ln(8)
    pdf.Cell(40, 10, "Location: "+resume.PersonalDetails.Location)

    pdf.Ln(12)
    pdf.SetFont(fontFamily, "B", fontSize+2)
    pdf.Cell(40, 10, "Professional Summary")
    pdf.Ln(8)
    pdf.SetFont(fontFamily, "", fontSize)
    pdf.MultiCell(0, 10, resume.ProfessionalSummary)

    return pdf.OutputFileAndClose(filename)
}