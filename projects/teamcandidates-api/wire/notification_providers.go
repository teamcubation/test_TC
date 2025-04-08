package wire

import (
	"errors"

	mdw "github.com/teamcubation/teamcandidates/pkg/http/middlewares/gin"
	gin "github.com/teamcubation/teamcandidates/pkg/http/servers/gin"
	smtp "github.com/teamcubation/teamcandidates/pkg/notification/smtp"

	notification "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/notification"
)

func ProvideNotificationSmtpService(smtp smtp.Service) (notification.SmtpService, error) {
	if smtp == nil {
		return nil, errors.New("smtp service cannot be nil")
	}
	return notification.NewSmtpService(smtp), nil
}

func ProvideNotificationUseCases(
	ssrv notification.SmtpService,
) notification.UseCases {
	return notification.NewUseCases(ssrv)
}

func ProvideNotificationHandler(
	server gin.Server,
	usecases notification.UseCases,
	middlewares *mdw.Middlewares,
) *notification.Handler {
	return notification.NewHandler(server, usecases, middlewares)
}
