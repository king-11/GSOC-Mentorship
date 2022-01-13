package mailer

import (
	"github.com/king-11/mentorship/extractor"
)

func(sender *Sender) MentorshipMail(mentorship extractor.BasicMentorship, template string) error {
  to := mentorship.GetEmails()
	body := sender.WriteHTMLEmail(to, "GSOC Mentorship Program - Mentor Allotment | COPS, IIT(BHU)", mentorship, template)

	return sender.SendMail(to, body)
}
