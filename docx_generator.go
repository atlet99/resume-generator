package main

import (
    "github.com/unidoc/unioffice/document"
    "os"
    "strconv"
)

func generateResumeDocx(resume Resume, filename string) error {
    loadEnv()

    doc := document.New()
    defer doc.Close()

    fontFamily := os.Getenv("FONT_FAMILY")
    headerFontSizeStr := os.Getenv("DOCX_HEADER_FONT_SIZE")
    headerFontSize, _ := strconv.Atoi(headerFontSizeStr)

    p := doc.AddParagraph()
    run := p.AddRun()
    run.AddText("Resume")
    run.Properties().SetSize(float64(headerFontSize)) // Преобразуем int в float64
    run.Properties().SetFontFamily(fontFamily)

    p = doc.AddParagraph()
    run = p.AddRun()
    run.AddText("Name: " + resume.PersonalDetails.Name)
    run.Properties().SetFontFamily(fontFamily)

    return doc.SaveToFile(filename)
}