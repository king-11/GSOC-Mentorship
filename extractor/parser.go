package extractor

import (
	"encoding/json"
	"io"
	"sort"
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
	sort.Slice(mentees, func(i, j int) bool {
		return mentees[i].Email < mentees[j].Email
	})
	return mentees, nil
}

func SetUpMentorship(mentees []*Mentee, mentors map[string]*Person) []*Mentorship {
	mentorships := make([]*Mentorship, 0)
	for name, val := range mentors {
		mentorship := &Mentorship{
			Mentor:  val,
			Mentees: []*Mentee{},
		}
		for _, mentee := range mentees {
			if mentee.Mentor == name {
				mentorship.Mentees = append(mentorship.Mentees, mentee)
			}
		}
		if len(mentorship.Mentees) > 0 {
			mentorships = append(mentorships, mentorship)
		}
	}

	return mentorships
}
