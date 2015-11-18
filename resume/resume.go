/*
Package resume implements the data model of a resume.
*/
package resume

func init() {
}

type Resume struct {
	name    string
	dob     string
	phone   string
	email   string
	addr    string
	website string
	intro   string
}
