package wire

import (
	"errors"

	gorm "github.com/teamcubation/teamcandidates/pkg/databases/sql/gorm"
	mdw "github.com/teamcubation/teamcandidates/pkg/http/middlewares/gin"
	ginsrv "github.com/teamcubation/teamcandidates/pkg/http/servers/gin"

	user "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/user"
)

func ProvideUserRepository(repo gorm.Repository) (user.Repository, error) {
	if repo == nil {
		return nil, errors.New("gorm repository cannot be nil")
	}
	return user.NewRepository(repo), nil
}

func ProvideUserUseCases(repo user.Repository) user.UseCases {
	return user.NewUseCases(repo)
}

func ProvideUserHandler(server ginsrv.Server, usecases user.UseCases, middlewares *mdw.Middlewares) *user.Handler {
	return user.NewHandler(server, usecases, middlewares)
}
