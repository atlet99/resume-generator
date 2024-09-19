package main

import (
    "github.com/unidoc/unioffice/document"
)

func generateResumeDocx(resume Resume, filename string) error {
    doc := document.New()
    defer doc.Close()

    doc.AddParagraph().AddRun().AddText("Resume")

    doc.AddParagraph().AddRun().AddText("Name: " + resume.PersonalDetails.Name)
    doc.AddParagraph().AddRun().AddText("Email: " + resume.PersonalDetails.Email)
    doc.AddParagraph().AddRun().AddText("Phone: " + resume.PersonalDetails.Phone)
    doc.AddParagraph().AddRun().AddText("Location: " + resume.PersonalDetails.Location)

    doc.AddParagraph().AddRun().AddText("Professional Summary")
    doc.AddParagraph().AddRun().AddText(resume.ProfessionalSummary)

    doc.AddParagraph().AddRun().AddText("Key Skills")
    for _, skill := range resume.KeySkills.Skills {
        doc.AddParagraph().AddRun().AddText("- " + skill)
    }

    return doc.SaveToFile(filename)
}