package excluder

import (
	"encoding/json"
	"io"

	"github.com/king-11/mentorship/extractor"
)

func GetMentees(f io.Reader) ([]string, error) {
	mentees := []*extractor.Mentee{}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&mentees); err != nil {
		return nil, err
	}

	emails := make([]string, 0)
	for _, val := range mentees {
		if val.Mentor != "" {
			emails = append(emails, val.Email)
		}
	}
	return emails, nil
}
