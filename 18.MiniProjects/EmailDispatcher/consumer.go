package main

import (
	"fmt"
	"net/smtp"
	"sync"
)

func emailWorker(id int, recipentChan <-chan Recipent, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipent := range recipentChan {
		smtpHost := "localhost"
		smtpPort := "1025"
		subject := fmt.Sprintf("Hello %s", recipent.Name)
		body := fmt.Sprintf("Hi %s,\n\nThis is a test email.", recipent.Name)
		msg := fmt.Sprintf("From: bilal@gmail.com\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", recipent.Email, subject, body)
		bytemsg := []byte(msg)
		smtp.SendMail(smtpHost+":"+smtpPort, nil, "bilal@gmail.com", []string{recipent.Email}, bytemsg)
		fmt.Println(id, recipent, msg)
	}
}
