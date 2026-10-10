package service

import "log"

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (s *EmailService) SendWelcomeEmail(email, fullName string) {
	// просто лог
	log.Printf("Welcome email sent to %s <%s>", fullName, email)

}
