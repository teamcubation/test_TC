package wire

import (
	"errors"

	gorm "github.com/teamcubation/teamcandidates/pkg/databases/sql/gorm"
	mdw "github.com/teamcubation/teamcandidates/pkg/http/middlewares/gin"
	ginsrv "github.com/teamcubation/teamcandidates/pkg/http/servers/gin"

	"github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/category"
)

func ProvideCategoryRepository(repo gorm.Repository) (category.Repository, error) {
	if repo == nil {
		return nil, errors.New("gorm repository cannot be nil")
	}
	return category.NewRepository(repo), nil
}

func ProvideCategoryUseCases(
	repo category.Repository,
) category.UseCases {
	return category.NewUseCases(repo)
}

func ProvideCategoryHandler(
	server ginsrv.Server,
	usecases category.UseCases,
	middlewares *mdw.Middlewares,
) *category.Handler {
	return category.NewHandler(server, usecases, middlewares)
}
