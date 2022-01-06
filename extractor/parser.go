package extractor

import (
	"encoding/json"
	"io"
)

func GetMentors(f io.Reader) (map[string]*Person, error) {
	mentors := make(map[string]*Person)
	decoder := json.NewDecoder(f)
	sliceMentors := []*Person{}
	if err := decoder.Decode(&sliceMentors); err != nil {
		return nil, err
	}
	for _, val := range sliceMentors {
		mentors[val.Name] = val
	}

	return mentors, nil
}

func GetMentees(f io.Reader) ([]*Mentee, error) {
	mentees := []*Mentee{}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&mentees); err != nil {
		return nil, err
	}
	for _, val := range mentees {
		val.Github = getGithubHandle(val.Github)
	}
	return mentees, nil
}

func SetUpMentorship(mentees []*Mentee, mentors map[string]*Person) []*Mentorship {
	mentorships := make([]*Mentorship, 0)
	for name, val := range mentors {
		mentorship := &Mentorship{
			Mentor: *val,
		}
		currentMentees := []Mentee{}
		for _, mentee := range mentees {
			if mentee.Mentor == name {
				currentMentees = append(currentMentees, *mentee)
			}
		}
		mentorship.Mentees = currentMentees
		if len(currentMentees) > 0 {
			mentorships = append(mentorships, mentorship)
		}
	}

	return mentorships
}
