/*
Package resume implements the data model of a resume.
*/
package resume

func init() {
}

type Resume struct {
	name            string
	dob             string
	phone           string
	email           string
	addr            string
	website         string
	intro           string
	summary         Summary
	workExperiences []WorkExperience
	projects        []Project
	edu             []Education
}

type Skill struct {
	sLevel string
	sType  string
	sName  string
}

type Summary struct {
	background string
	ethic      string
	skills     []Skill
}

type WorkExperience struct {
	title    string
	company  string
	start    string
	end      string
	location string
	projects []Project
}

type Project struct {
	name    string
	start   string
	end     string
	desc    string
	works   []string
	contrib string
	teches  []string
}

type Education struct {
	school      string
	location    string
	start       string
	end         string
	degreeLevel string
	major       string
	gpa         float32
	gpaMax      float32
}
