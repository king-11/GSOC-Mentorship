package mailer

import (
	"github.com/king-11/mentorship/extractor"
)

func(sender *Sender) MentorshipMail(mentorship *extractor.Mentorship) error {
  to := mentorship.GetEmails()
	body := sender.WriteHTMLEmail(to, "GSOC Mentorship Program - Mentor Allotment | COPS, IIT(BHU)", mentorship)

	return sender.SendMail(to, body)
}
