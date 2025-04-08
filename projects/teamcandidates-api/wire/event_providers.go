package wire

import (
	"errors"

	mng "github.com/teamcubation/teamcandidates/pkg/databases/nosql/mongodb/mongo-driver"
	mdw "github.com/teamcubation/teamcandidates/pkg/http/middlewares/gin"
	ginsrv "github.com/teamcubation/teamcandidates/pkg/http/servers/gin"

	event "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/event"
)

func ProvideEventRepository(repo mng.Repository) (event.Repository, error) {
	if repo == nil {
		return nil, errors.New("mongoDB repository cannot be nil")
	}
	return event.NewRepository(repo), nil
}

func ProvideEventUseCases(repo event.Repository) event.UseCases {
	return event.NewUseCases(repo)
}

func ProvideEventHandler(server ginsrv.Server, usecases event.UseCases, middlewares *mdw.Middlewares) *event.Handler {
	return event.NewHandler(server, usecases, middlewares)
}
