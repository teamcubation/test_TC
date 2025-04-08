package wire

import (
	"errors"

	gorm "github.com/teamcubation/teamcandidates/pkg/databases/sql/gorm"
	mdw "github.com/teamcubation/teamcandidates/pkg/http/middlewares/gin"
	ginsrv "github.com/teamcubation/teamcandidates/pkg/http/servers/gin"

	candidate "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/candidate"
)

// ProvideCandidateRepository retorna candidate.Repository a partir de un repositorio GORM.
func ProvideCandidateRepository(repo gorm.Repository) (candidate.Repository, error) {
	if repo == nil {
		return nil, errors.New("gorm repository cannot be nil")
	}
	return candidate.NewRepository(repo), nil
}

// ProvideCandidateUseCases retorna candidate.UseCases a partir del repositorio.
func ProvideCandidateUseCases(repo candidate.Repository) candidate.UseCases {
	return candidate.NewUseCases(repo)
}

// ProvideCandidateHandler retorna el Handler de candidate inyectando el servidor Gin,
// el servidor WebSocket, los casos de uso y los middlewares.
func ProvideCandidateHandler(
	ginSrv ginsrv.Server,
	usecases candidate.UseCases,
	middlewares *mdw.Middlewares,
) *candidate.Handler {
	return candidate.NewHandler(ginSrv, usecases, middlewares)
}
