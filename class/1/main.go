package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailersend/mailersend-go"
)

func main() {
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend("mlsn.34598713dfd7d71e7dfc181a26b904c029cbdae88595a27e42eee94f1b426811")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	subject := "Subject"
	text := "This is the text content"
	html := "<p>This is the HTML content</p>"

	from := mailersend.From{
		Name:  "Toko Mantap",
		Email: "MS_q1mJQr@trial-351ndgw0nrd4zqx8.mlsender.net",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "Paul",
			Email: "paulus.otto.harman@gmail.com",
		},
	}

	// Send in 5 minute
	//sendAt := time.Now().Add(time.Minute * 5).Unix()

	//tags := []string{"foo", "bar"}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	//message.SetTags(tags)
	//message.SetSendAt(sendAt)
	//message.SetInReplyTo("client-id")

	res, err := ms.Email.Send(ctx, message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(res.Header.Get("X-Message-Id"))

}
