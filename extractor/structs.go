package extractor

import (
	"fmt"
	"strings"
)

type Person struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type BasicMentorship interface {
	GetEmails() []string
	String() string
	GetMentorNames() string
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

type MutliMentorship struct {
	Mentor  []*Person
	Mentees []*Mentee
	ID int
}

func(m *MutliMentorship) GetMentorNames() string {
	if len(m.Mentor) == 1 {
		return m.Mentor[0].Name
	}
	names := make([]string, len(m.Mentor) - 1)
	for i, mentor := range m.Mentor {
		if i == len(m.Mentor) - 1 {
			break
		}
		names[i] = mentor.Name
	}
	return fmt.Sprintf("%s and %s", strings.Join(names, ", "), m.Mentor[len(m.Mentor) - 1].Name)
}

func(m *MutliMentorship) GetEmails() []string {
	emails := []string{}
	for _, mentee := range m.Mentees {
		emails = append(emails, mentee.Email)
	}
	for _, mentor := range m.Mentor {
		emails = append(emails, mentor.Email)
	}
	return emails
}

func(m *MutliMentorship) String() string {
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
