package mail

import (
	"emailn/internal/domain/campaign"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(campaign *campaign.Campaign) error {
    m := gomail.NewMessage()

    port, err := strconv.ParseInt(os.Getenv("EMAIL_PORT"), 10, 32)

    if err != nil {
        log.Fatal("Invalid Email Port", err.Error())
    }

    d := gomail.NewDialer(os.Getenv("EMAIL_SMTP_PROVIDER"), int(port), os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

    var emails []string

    for _, contact := range campaign.Contacts {
        emails = append(emails, contact.Email)
    }

    m.SetHeader("From", os.Getenv("EMAIL_USER"))
    m.SetHeader("To", emails...)
    m.SetHeader("Subject", campaign.Name)
    m.SetBody("text/html", campaign.Content)

    return d.DialAndSend(m)
}
