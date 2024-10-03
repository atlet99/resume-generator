package main

import (
	"fmt"
)

func main() {
	resume := Resume{
		PersonalDetails: PersonalDetails{
			Name:    "John Doe",
			Email:   "john.doe@example.com",
			Phone:   "+44 1234 567890",
			Address: "123 Baker Street, London, UK",
		},
		ProfessionalSummary: "Experienced software engineer with over 5 years in backend development...",
		KeySkills:           []string{"Go", "Docker", "Kubernetes", "AWS", "CI/CD"},
		WorkExperience: []WorkExperience{
			{
				JobTitle:    "Senior Engineer",
				Company:     "TechCorp",
				StartDate:   "Jan 2020",
				EndDate:     "Present",
				Description: "Led backend development using Go and microservices.",
			},
		},
		Education: []Education{
			{
				Degree:     "BSc Computer Science",
				University: "University of London",
				Year:       "2017",
			},
		},
		Certifications: []Certification{
			{
				Title: "AWS Certified Solutions Architect",
				Year:  "2019",
			},
		},
		Languages: []string{"English", "Spanish"},
	}

	docxFilename := fmt.Sprintf("%s.docx", resume.PersonalDetails.Name)
	pdfFilename := fmt.Sprintf("%s.pdf", resume.PersonalDetails.Name)

	err := generateResumeDocx(resume, docxFilename)
	if err != nil {
		panic(err)
	}

	err = generateResumePDF(resume, pdfFilename)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Resume generated successfully! Files saved to %s and %s\n", docxFilename, pdfFilename)
}
