package dto

import (
	"github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/notification/usecases/domain"
)

type Email struct {
	Address string `json:"address"`
	Subject string `json:"subject"`
	Body    string `json:"body_template"`
}

// ToDomain convierte un Email en una entidad del dominio (domain.Email).
func (e *Email) ToDomain() *domain.Email {
	return &domain.Email{
		Address: e.Address,
		Subject: e.Subject,
		Body:    e.Body,
	}
}

type VerificationMessage struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type Request struct {
	Email string `json:"email" binding:"required,email"`
}
