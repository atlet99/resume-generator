package main

type Resume struct {
    PersonalDetails      PersonalDetails
    ProfessionalSummary  string
    KeySkills            []string
    WorkExperience       []WorkExperience
    Education            []Education
    Certifications       []Certification
    Languages            []string
}

type PersonalDetails struct {
    Name     string
    Email    string
    Phone    string
    Address  string
}

type WorkExperience struct {
    JobTitle    string
    Company     string
    StartDate   string
    EndDate     string
    Description string
}

type Education struct {
    Degree     string
    University string
    Year       string
}

type Certification struct {
    Title string
    Year  string
}