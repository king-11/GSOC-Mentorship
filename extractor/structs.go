package extractor

import "fmt"

type Person struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type Mentorship struct {
	Mentor  *Person
	Mentees []*Mentee
}

func(m *Mentorship) GetMentorNames() string {
	return m.Mentor.Name
}

func(m *Mentorship) GetEmails() []string {
	emails := []string{}
	for _, mentee := range m.Mentees {
		emails = append(emails, mentee.Email)
	}
	emails = append(emails, m.Mentor.Email)
	return emails
}

func(m *Mentorship) String() string {
	return fmt.Sprintf("%s:%+v",m.GetMentorNames(), m.GetEmails())
}

type Mentee struct {
	Year       string `json:"Year"`
	Experience string `json:"Experience"`
	Name       string `json:"Name"`
	Email      string `json:"Email"`
	Mentor     string `json:"Mentor"`
	Github     string `json:"Github Handle"`
	Contact    int    `json:"Phone Number"`
}
