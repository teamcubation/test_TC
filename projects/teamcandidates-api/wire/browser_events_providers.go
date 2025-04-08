package wire

import (
	"errors"

	mng "github.com/teamcubation/teamcandidates/pkg/databases/nosql/mongodb/mongo-driver"
	mdw "github.com/teamcubation/teamcandidates/pkg/http/middlewares/gin"
	ginsrv "github.com/teamcubation/teamcandidates/pkg/http/servers/gin"
	ws "github.com/teamcubation/teamcandidates/pkg/websocket/gorilla"

	browserevent "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/browser-events"
)

func ProvideBrowserEventsRepository(repo mng.Repository) (browserevent.Repository, error) {
	if repo == nil {
		return nil, errors.New("mongoDB repository cannot be nil")
	}
	return browserevent.NewRepository(repo), nil
}

// ProvideBrowserEventsUseCases retorna browserevent.UseCases a partir del repositorio.
func ProvideBrowserEventsUseCases(repo browserevent.Repository) browserevent.UseCases {
	return browserevent.NewUseCases(repo)
}

func ProvideBrowserEventsWebsocket(
	useCases browserevent.UseCases,
	upgrader ws.Upgrader,
) browserevent.WebSocket {
	return browserevent.NewWebSocket(useCases, upgrader)
}

// ProvideBrowserEventsHandler retorna el Handler de browserevent inyectando el servidor Gin,
// el servidor WebSocket, los casos de uso y los middlewares.
func ProvideBrowserEventsHandler(
	ginSrv ginsrv.Server,
	usecases browserevent.UseCases,
	middlewares *mdw.Middlewares,
	websocket browserevent.WebSocket,
) *browserevent.Handler {
	return browserevent.NewHandler(ginSrv, usecases, middlewares, websocket)
}
