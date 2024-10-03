package utils

import (
	"fmt"
	"resume-generator/models"
)

func ValidateMandatoryFields(resume models.Resume) error {
	if resume.Name == nil || *resume.Name == "" {
		return fmt.Errorf("Name is mandatory and must be provided")
	}
	if resume.Phone == nil || *resume.Phone == "" {
		return fmt.Errorf("Phone is mandatory and must be provided")
	}
	if resume.Email == nil || *resume.Email == "" {
		return fmt.Errorf("Email is mandatory and must be provided")
	}
	return nil
}
