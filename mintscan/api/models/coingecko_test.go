package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

var data = []byte(`{
  "body": {
    "template": "abctemplate.html",
    "params": {
      "name": "Chase",
      "email": "1234@gmail.com"
    }
  },
  "to": [
	  [
    	"abc@gmail.com",
    	"xyz@gmail.com"
	  ],
	  [
    	"jayb@gmail.com",
    	"kim@gmail.com"
	  ]
  ],
  "cc": [
    "xxx@example.com",
    "yyy@example.com"
  ],
  "replyTo": {
    "email": "aaa@gmail.com",
    "name": "Jack"
  }
}`)

type emailData struct {
	Body struct {
		Template string            `json:"template"`
		Params   map[string]string `json:"params"`
	} `json:"body"`
	To      [][]string `json:"to"`
	CC      []string   `json:"cc"`
	ReplyTo struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
}

func TestPayload(t *testing.T) {
	var email emailData
	if err := json.Unmarshal(data, &email); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", email.To)
}
