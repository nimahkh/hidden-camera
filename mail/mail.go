package mail

import (
	"fmt"
	"github.com/scorredoira/email"
	"net/mail"
	"net/smtp"
	"os"
)

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func Send() {
	pwd, _ := os.Getwd()
	attachment := pwd + "/save.png"

	from := "nima.2004hkh@gmail.com" // change here
	pass := "manimani226688"
	to := []string{"nima.2004hkh@gmail.com"}
	message := "Please see the email attachment for the images"
	subject := "Attached my photos!"
	msg := email.NewMessage(subject, message)
	msg.From = mail.Address{Name: "From", Address: from}
	msg.To = to

	//read file
	err := msg.Attach(attachment)
	if err != nil {
		fmt.Println(err)
	}

	err = email.Send("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		msg)

	if err != nil {
		fmt.Println(err)
	}
}
