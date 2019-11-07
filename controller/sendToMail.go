package controller

import (
	"fmt"
	"net/smtp"
	"strings"
)

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func SendMail(Body, Subject, to string) bool {
	err := SendToMail(EMAILFROM, EMAILPASSWORD, EMAILHOST, to, Subject, Body, "html")
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

func MailTemplate(code, to string) bool {
	var Body = `
		<html>
			<body>
				<h3>
					` + "genaro xpool verification code: " + code + `
				</h3>
			</body>
		</html>
		`
	Subject := "genaro xpool"
	return SendMail(Body, Subject, to)
}
