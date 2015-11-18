/*
Package resume implements the data model of a resume.
*/
package resume

func init() {
}

type Resume struct {
	Name            string
	Dob             string
	Phone           string
	Email           string
	Addr            string
	Website         string
	Intro           string
	Summary         Summary
	WorkExperiences []WorkExperience
	Projects        []Project
	Edu             []Education
}

type Skill struct {
	SLevel string
	SType  string
	SName  string
}

type Summary struct {
	Background string
	Ethic      string
	Skills     []Skill
}

type WorkExperience struct {
	Title    string
	Company  string
	Start    string
	End      string
	Location string
	Projects []Project
}

type Project struct {
	Name    string
	Start   string
	End     string
	Desc    string
	Works   []string
	Contrib string
	Teches  []string
}

type Education struct {
	School      string
	Location    string
	Start       string
	End         string
	DegreeLevel string
	Major       string
	Gpa         float32
	GpaMax      float32
}
