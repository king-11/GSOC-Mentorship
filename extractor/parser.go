package extractor

import (
	"encoding/json"
	"io"
	"sort"
	"strings"
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

func SetupMultiMentorship(mentees []*Mentee, mentorsMap map[string]*Person) []*MutliMentorship {
	mentorshipsMap := make(map[string]*MutliMentorship)
	for _, val := range mentees {
		if ment, ok := mentorshipsMap[val.Mentor]; ok {
			ment.Mentees = append(ment.Mentees, val)
		} else {
			mentorshipsMap[val.Mentor] = &MutliMentorship{
				Mentor:  []*Person{},
				Mentees: []*Mentee{val},
			}
			mentors := strings.Split(val.Mentor, ",")
			for _, mentor := range mentors {
				mentor := strings.Trim(mentor, " ")
				if _, ok := mentorsMap[mentor]; ok {
					mentorshipsMap[val.Mentor].Mentor = append(mentorshipsMap[val.Mentor].Mentor, mentorsMap[mentor])
				}
			}
		}
	}

	mentorships := make([]*MutliMentorship, 0)
	for _, val := range mentorshipsMap {
		mentorships = append(mentorships, val)
	}

	sort.Slice(mentorships, func(i, j int) bool {
		return mentorships[i].Mentor[0].Name < mentorships[j].Mentor[0].Name
	})
	for idx, val := range mentorships {
		val.ID = idx + 1
	}
	return mentorships
}
