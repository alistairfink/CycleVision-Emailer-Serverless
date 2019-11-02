package main

import (
	"github.com/alistairfink/iBlind-Emailer-Serverless/p"
)

func main() {
	request := p.Request{
		Name:    "Alistair Fink",
		Email:   "alistairfink@gmail.com",
		Subject: "Test Message",
		Message: "Hello World!",
	}

	if p.Logic(request) {
		println("Message Successfully Sent!")
	} else {
		println("Message Failed to Send!")
	}
}
