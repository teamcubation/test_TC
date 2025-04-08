package wire

import (
	"errors"

	gorm "github.com/teamcubation/teamcandidates/pkg/databases/sql/gorm"
	mdw "github.com/teamcubation/teamcandidates/pkg/http/middlewares/gin"
	ginsrv "github.com/teamcubation/teamcandidates/pkg/http/servers/gin"

	"github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/assessment"
	"github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/authe"
	"github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/candidate"
	"github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/config"
	"github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/notification"
	"github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/person"
)

// ProvideAssessmentRepository inyecta la implementación de Repository para Assessment.
func ProvideAssessmentRepository(repo gorm.Repository) (assessment.Repository, error) {
	if repo == nil {
		return nil, errors.New("gorm repository cannot be nil")
	}
	return assessment.NewRepository(repo), nil
}

// ProvideAssessmentUseCases inyecta las dependencias requeridas por la capa de casos de uso de Assessment.
func ProvideAssessmentUseCases(
	repo assessment.Repository,
	notif notification.UseCases,
	cand candidate.UseCases,
	cfg config.Loader,
	au authe.UseCases,
	pn person.UseCases,
) assessment.UseCases {
	return assessment.NewUseCases(repo, notif, cand, cfg, au, pn)
}

// ProvideAssessmentHandler inyecta las dependencias para crear el Handler de Assessment.
func ProvideAssessmentHandler(
	server ginsrv.Server,
	usecases assessment.UseCases,
	middlewares *mdw.Middlewares,
) *assessment.Handler {
	return assessment.NewHandler(server, usecases, middlewares)
}
