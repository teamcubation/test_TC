package wire

import (
	"errors"

	gorm "github.com/teamcubation/teamcandidates/pkg/databases/sql/gorm"
	mdw "github.com/teamcubation/teamcandidates/pkg/http/middlewares/gin"
	ginsrv "github.com/teamcubation/teamcandidates/pkg/http/servers/gin"

	group "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/group"
)

func ProvideGroupRepository(repo gorm.Repository) (group.Repository, error) {
	if repo == nil {
		return nil, errors.New("gorm repository cannot be nil")
	}
	return group.NewRepository(repo), nil
}

func ProvideGroupUseCases(repo group.Repository) group.UseCases {
	return group.NewUseCases(repo)
}

func ProvideGroupHandler(server ginsrv.Server, usecases group.UseCases, middlewares *mdw.Middlewares) *group.Handler {
	return group.NewHandler(server, usecases, middlewares)
}
