package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/king-11/mentorship/excluder"
	"github.com/king-11/mentorship/extractor"
	"github.com/king-11/mentorship/mailer"
)

const (
	// smtp user config
	FROM = "iitbhu.cops@gmail.com"
)

func main() {
  f, err := os.Open("mentors.json")
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()
  mentors, err := extractor.GetMentors(f)
  if err != nil {
    log.Fatal(err)
  }

  fl, err := os.Open("mentees.json")
  if err != nil {
    log.Fatal(err)
  }
  defer fl.Close()
  mentees_all, err := extractor.GetMentees(fl)
  if err != nil {
    log.Fatal(err)
  }

  fe, err := os.Open("mentees_1.json")
  if err != nil {
    log.Fatal(err)
  }
  defer fe.Close()
  mentees_done, err := excluder.GetMentees(fe)
  if err != nil {
    log.Fatal(err)
  }


  mentees := make([]*extractor.Mentee, 0)
  for _, val := range mentees_all {
    check := false
    if strings.Contains(val.Email, ",") {
      continue
    }
    for _, v := range mentees_done {
      if val.Email == v {
        check = true
        break
      }
    }

    if !check {
      mentees = append(mentees, val)
    }
  }
  mentorships := extractor.SetUpMentorship(mentees, mentors)

  err = godotenv.Load(".env")
  if err != nil {
    log.Fatal(err)
  }
  // google app password
  PASSWORD := os.Getenv("PASSWORD")

  sender := mailer.NewSender(FROM, PASSWORD)
  for _, val :=  range mentorships {
    // write html email
    // body := sender.WriteHTMLEmail(val.GetEmails(), "GSOC", val)
    // tempFile, err := os.Create("htmls/" + val.Mentor.Name + ".html")
    // if err != nil {
    //   log.Fatal(err)
    // }
    // defer tempFile.Close()
    // tempFile.Write(body)

    // send mail
    err := sender.MentorshipMail(val)
    if err != nil {
      log.Fatal(err)
    }
    log.Printf("Sent mail to group %v\n", val)
  }
}
