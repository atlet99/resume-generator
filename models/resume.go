package models

type Experience struct {
	Role        *string
	Company     *string
	Location    *string
	Dates       *string
	Description *string
}

type Resume struct {
	Name            *string
	Phone           *string
	Email           *string
	LinkedIn        *string
	GitHub          *string
	TechnicalSkills *string
	Experience      []Experience
	Education       *string
}
