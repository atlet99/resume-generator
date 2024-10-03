package main

import (
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/atlet99/resume-generator/env"
	"github.com/atlet99/resume-generator/formats"
	"github.com/atlet99/resume-generator/models"
	"github.com/atlet99/resume-generator/utils"
)

func main() {
	err := env.LoadEnv(".template.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	experience, err := env.GetExperience()
	if err != nil {
		log.Fatalf("Error loading experience: %v", err)
	}

	resume := models.Resume{
		Name:            env.GetEnv("NAME"),
		Phone:           env.GetEnv("PHONE"),
		Email:           env.GetEnv("EMAIL"),
		LinkedIn:        env.GetEnv("LINKEDIN"),
		GitHub:          env.GetEnv("GITHUB"),
		TechnicalSkills: env.GetEnv("TECHNICAL_SKILLS"),
		Experience:      experience,
		Education:       env.GetEnv("EDUCATION"),
	}

	err = utils.ValidateMandatoryFields(resume)
	if err != nil {
		log.Fatalf("Validation error: %v", err)
	}

	format, err := extractFormatFromTemplate("templates/resume.template")
	if err != nil {
		log.Fatalf("Error extracting format from template: %v", err)
	}

	if format == "" {
		format = "BOTH"
	}

	switch strings.ToUpper(format) {
	case "PDF":
		generateResumeInFormat(resume, "pdf", "generated_resume.pdf")
	case "DOC":
		generateResumeInFormat(resume, "docx", "generated_resume.docx")
	case "BOTH":
		generateResumeInFormat(resume, "pdf", "generated_resume.pdf")
		generateResumeInFormat(resume, "docx", "generated_resume.docx")
	default:
		log.Fatalf("Invalid format specified in template: %v", format)
	}
}

func generateResumeInFormat(resume models.Resume, format string, filename string) {
	var err error
	switch format {
	case "pdf":
		err = formats.GeneratePDF(resume, filename)
	case "docx":
		err = formats.GenerateDOC(resume, filename)
	default:
		log.Fatalf("Invalid format: %v", format)
	}

	if err != nil {
		log.Fatalf("Error generating %s: %v", format, err)
	}

	log.Printf("Resume generated successfully in %s format.\n", format)
}

func extractFormatFromTemplate(templateFile string) (string, error) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}

	data := struct {
		Format string
	}{
		Format: "",
	}

	err = tmpl.Execute(os.Stdout, &data)
	if err != nil {
		return "", err
	}

	return data.Format, nil
}
