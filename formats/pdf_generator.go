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

	err = pdf.Cell(nil, "Resume")
	if err != nil {
		log.Print(err.Error())
		return err
	}
	pdf.Br(20)

	if resume.Name != nil {
		err = pdf.Cell(nil, "Name: "+*resume.Name)
		if err != nil {
			log.Print(err.Error())
			return err
		}
		pdf.Br(20)
	}

	if resume.Phone != nil {
		err = pdf.Cell(nil, "Phone: "+*resume.Phone)
		if err != nil {
			log.Print(err.Error())
			return err
		}
		pdf.Br(20)
	}

	if resume.Email != nil {
		err = pdf.Cell(nil, "Email: "+*resume.Email)
		if err != nil {
			log.Print(err.Error())
			return err
		}
		pdf.Br(20)
	}

	if len(resume.Experience) > 0 {
		err = pdf.Cell(nil, "Experience:")
		if err != nil {
			log.Print(err.Error())
			return err
		}
		pdf.Br(20)
		for _, exp := range resume.Experience {
			if exp.Role != nil {
				err = pdf.Cell(nil, "Role: "+*exp.Role)
				if err != nil {
					log.Print(err.Error())
					return err
				}
				pdf.Br(15)
			}
			if exp.Company != nil {
				err = pdf.Cell(nil, "Company: "+*exp.Company)
				if err != nil {
					log.Print(err.Error())
					return err
				}
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
