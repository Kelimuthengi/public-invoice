package services

import (
	"log"

	"github.com/wneessen/go-mail"
)

type EmailMessage struct {
	RecepientEmail string
	Subject        string
}

// func (em *EmailMessage) SendMailData(emailData ) error {

// 	m := mail.NewMsg()

// 	if err := m.From(senderemail); err != nil {
// 		log.Fatalf("Provide valid sender address %s", err)
// 	}

// 	if err := m.To(receiptemail); err != nil {
// 		log.Fatalf("Provide valid receiver address%s", err)
// 	}
// 	return nil
// }

func SendMail(emailData EmailMessage) error {

	m := mail.NewMsg()

	if err := m.From("payday" + " <" + "theeoligarchy@gmail.com" + ">"); err != nil {
		log.Fatalf("Provide valid sender address %s", err)
	}

	if err := m.To(emailData.RecepientEmail); err != nil {
		log.Fatalf("Provide valid receiver address%s", err)
	}

	m.Subject("New_Dawn Apartments")
	m.SetBodyString(mail.TypeTextPlain, emailData.Subject)

	data, err := mail.NewClient("smtp.gmail.com", mail.WithPort(25), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername("theeoligarchy@gmail.com"), mail.WithPassword("qble klei ymrp zzpv"))
	if err != nil {

		return err
	}

	if err := data.DialAndSend(m); err != nil {

		return err
	}

	return nil
}
