package main

import (
	"github.com/alistairfink/iBlind-Emailer-Serverless/p"
)

func main() {
	request := p.Request{
		Name:    "Alistair Fink",
		Email:   "alistairfink@gmail.com",
		Message: "Hello World!",
	}

	if p.Logic(request) {
		println("Successfully Sent!")
	} else {
		println("Message Failed to Send!")
	}
}
