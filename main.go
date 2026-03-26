package main

import "fmt"

type Notification interface {
	Send()
}

type EmailService struct{}

func (e EmailService) Send() {
	fmt.Println("email send")
}

type SMSService struct{}

func (s SMSService) Send() {
	fmt.Println("SMS send")
}

type UserService struct {
	notifier1 Notification
	notifier2 Notification
}

func (u UserService) RegisterUser() {
	u.notifier1.Send()
	u.notifier2.Send()
}

func main() {
	var email Notification
	email = EmailService{}
	var sms Notification
	sms = SMSService{}

	var user UserService
	user = UserService{
		notifier1: email,
		notifier2: sms,
	}
	user.RegisterUser()

}
