package env

import (
	"fmt"
	"os"
	"resume-generator/models"

	"github.com/joho/godotenv"
)

func LoadEnv(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func GetEnv(key string) *string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return nil
	}
	if value == "" {
		return nil
	}
	return &value
}

func GetExperience() ([]models.Experience, error) {
	var experiences []models.Experience

	for i := 1; ; i++ {
		role := GetEnv(fmt.Sprintf("EXPERIENCE_%d_ROLE", i))
		if role == nil {
			break
		}

		experience := models.Experience{
			Role:        role,
			Company:     GetEnv(fmt.Sprintf("EXPERIENCE_%d_COMPANY", i)),
			Location:    GetEnv(fmt.Sprintf("EXPERIENCE_%d_LOCATION", i)),
			Dates:       GetEnv(fmt.Sprintf("EXPERIENCE_%d_DATES", i)),
			Description: GetEnv(fmt.Sprintf("EXPERIENCE_%d_DESCRIPTION", i)),
		}

		if experience.Role == nil || experience.Company == nil || experience.Location == nil || experience.Dates == nil || experience.Description == nil {
			return nil, fmt.Errorf("some of experience fields are empty for experience %d", i)
		}

		experiences = append(experiences, experience)
	}

	return experiences, nil
}
