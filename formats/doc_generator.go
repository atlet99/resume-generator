package formats

import (
	"log"
	"resume-generator/models"

	"github.com/unidoc/unioffice/document"
)

func GenerateDOC(resume models.Resume, filename string) error {
	doc := document.New()

	heading := doc.AddParagraph()
	run := heading.AddRun()
	run.AddText("Resume")
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("Arial")
	run.Properties().SetSize(14)

	if resume.Name != nil {
		paragraph := doc.AddParagraph()
		run := paragraph.AddRun()
		run.AddText("Name: " + *resume.Name)
		run.Properties().SetFontFamily("Arial")
		run.Properties().SetSize(12)
	}
	if resume.Phone != nil {
		paragraph := doc.AddParagraph()
		run := paragraph.AddRun()
		run.AddText("Phone: " + *resume.Phone)
		run.Properties().SetFontFamily("Arial")
		run.Properties().SetSize(12)
	}
	if resume.Email != nil {
		paragraph := doc.AddParagraph()
		run := paragraph.AddRun()
		run.AddText("Email: " + *resume.Email)
		run.Properties().SetFontFamily("Arial")
		run.Properties().SetSize(12)
	}

	if len(resume.Experience) > 0 {
		heading = doc.AddParagraph()
		run = heading.AddRun()
		run.AddText("Experience:")
		run.Properties().SetBold(true)
		run.Properties().SetFontFamily("Arial")
		run.Properties().SetSize(14)

		for _, exp := range resume.Experience {
			if exp.Role != nil {
				paragraph := doc.AddParagraph()
				run := paragraph.AddRun()
				run.AddText("Role: " + *exp.Role)
				run.Properties().SetFontFamily("Arial")
				run.Properties().SetSize(12)
			}
			if exp.Company != nil {
				paragraph := doc.AddParagraph()
				run := paragraph.AddRun()
				run.AddText("Company: " + *exp.Company)
				run.Properties().SetFontFamily("Arial")
				run.Properties().SetSize(12)
			}
		}
	}

	err := doc.SaveToFile(filename)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	return nil
}
