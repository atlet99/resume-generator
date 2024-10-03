package formats

import (
	"log"
	"resume-generator/models"

	"github.com/signintech/gopdf"
)

func GeneratePDF(resume models.Resume, filename string) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	err := pdf.AddTTFFont("Arial", "./assets/fonts/Arial.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.SetFont("Arial", "", 14)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	pdf.Cell(nil, "Resume")
	pdf.Br(20)

	if resume.Name != nil {
		pdf.Cell(nil, "Name: "+*resume.Name)
		pdf.Br(20)
	}
	if resume.Phone != nil {
		pdf.Cell(nil, "Phone: "+*resume.Phone)
		pdf.Br(20)
	}
	if resume.Email != nil {
		pdf.Cell(nil, "Email: "+*resume.Email)
		pdf.Br(20)
	}

	if len(resume.Experience) > 0 {
		pdf.Cell(nil, "Experience:")
		pdf.Br(20)
		for _, exp := range resume.Experience {
			if exp.Role != nil {
				pdf.Cell(nil, "Role: "+*exp.Role)
				pdf.Br(15)
			}
			if exp.Company != nil {
				pdf.Cell(nil, "Company: "+*exp.Company)
				pdf.Br(15)
			}
		}
	}

	err = pdf.WritePdf(filename)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	return nil
}
