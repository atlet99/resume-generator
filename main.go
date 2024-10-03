package main

import (
	"log"
	"resume-generator/env"
	"resume-generator/formats"
	"resume-generator/models"
	"resume-generator/utils"
)

func main() {
	// Загружаем .env файл
	err := env.LoadEnv(".template.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Собираем данные для резюме
	resume := models.Resume{
		Name:            env.GetEnv("NAME"),
		Phone:           env.GetEnv("PHONE"),
		Email:           env.GetEnv("EMAIL"),
		LinkedIn:        env.GetEnv("LINKEDIN"),
		GitHub:          env.GetEnv("GITHUB"),
		TechnicalSkills: env.GetEnv("TECHNICAL_SKILLS"),
		Experience:      env.GetExperience(),
		Education:       env.GetEnv("EDUCATION"),
	}

	err = utils.ValidateMandatoryFields(resume)
	if err != nil {
		log.Fatalf("Validation error: %v", err)
	}

	err = formats.GeneratePDF(resume, "generated_resume.pdf")
	if err != nil {
		log.Fatalf("Error generating PDF: %v", err)
	}

	err = formats.GenerateDOC(resume, "generated_resume.docx")
	if err != nil {
		log.Fatalf("Error generating DOC: %v", err)
	}

	log.Println("Resume generated successfully in both PDF and DOC formats.")
}
