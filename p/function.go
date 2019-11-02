package p

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Emailer(w http.ResponseWriter, r *http.Request) {
	request := Request{}
	if err := json.NewDecoder(r.Body).Decode(&request); err == nil {
		fmt.Fprint(w, "Message Failed to Send!")
		return
	}

	if Logic(request) {
		fmt.Fprint(w, "Successfully Sent!")
	} else {
		fmt.Fprint(w, "Message Failed to Send!")
	}
}

func Logic(request Request) bool {
	message := "Name: " + request.Name + "\n" +
		"Email: " + request.Email + "\n" +
		"Message: " + request.Message

	println(message)
	return true
}

type Request struct {
	Name    string `json: "name"`
	Email   string `json: "email"`
	Message string `json: "message"`
}
