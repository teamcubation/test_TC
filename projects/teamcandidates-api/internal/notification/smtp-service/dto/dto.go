package dto

import (
	"errors"

	smtp "github.com/teamcubation/teamcandidates/pkg/notification/smtp"

	"github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/notification/usecases/domain"
)

type Email struct {
	Address string
	Subject string
	Body    string
}

func FromDomain(email *domain.Email) (*smtp.Email, error) {
	if email == nil {
		return nil, errors.New("email cannot be nil")
	}
	return &smtp.Email{
		Address: email.Address,
		Subject: email.Subject,
		Body:    email.Body,
	}, nil
}
