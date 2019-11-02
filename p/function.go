package p

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

func Emailer(w http.ResponseWriter, r *http.Request) {
	request := Request{}
	if err := json.NewDecoder(r.Body).Decode(&request); err == nil {
		fmt.Fprint(w, "Message Failed to Send!")
		return
	}

	if Logic(request) {
		fmt.Fprint(w, "Message Successfully Sent!")
	} else {
		fmt.Fprint(w, "Message Failed to Send!")
	}
}

func Logic(request Request) bool {
	message := "Name: " + request.Name + "\n" +
		"Email: " + request.Email + "\n" +
		"Message: " + request.Message

	from := "alistairfinkraspberrypi@gmail.com"
	pass := ""
	to := "alistairfink@gmail.com"
	body := message
	subject := request.Subject

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return false
	}

	return true
}

type Request struct {
	Name    string `json: "name"`
	Email   string `json: "email"`
	Subject string `json: "subject"`
	Message string `json: "message"`
}
