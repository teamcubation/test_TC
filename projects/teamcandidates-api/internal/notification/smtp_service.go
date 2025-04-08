package notification

import (
	"context"

	smtp "github.com/teamcubation/teamcandidates/pkg/notification/smtp"

	dto "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/notification/smtp-service/dto"
	"github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/notification/usecases/domain"
)

type smtpService struct {
	smtpService smtp.Service
}

func NewSmtpService(ss smtp.Service) SmtpService {
	return &smtpService{
		smtpService: ss,
	}
}

func (ss *smtpService) SendEmail(ctx context.Context, data *domain.Email) error {
	formatedData, err := dto.FromDomain(data)
	if err != nil {
		return err
	}
	return ss.smtpService.SendEmail(ctx, formatedData)
}
