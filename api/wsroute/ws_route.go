package wsroute

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/hrishi-backend/api/routeutils"
	"gitlab.com/Aubichol/hrishi-backend/cfg"
	"gitlab.com/Aubichol/hrishi-backend/ws"
)

// TODO: Need to know what websocket upgrader does
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//wsHandler holds web socket handler
type wsHandler struct {
	hub       ws.Hub
	clientCfg cfg.WSClient
}

//ServeHTTP implements http.Handler interface for web scoket handler
func (wh *wsHandler) ServeHTTP(
	w http.ResponseWriter, r *http.Request,
) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Error(err)
		return
	}

	client := ws.NewDefaultClient(conn, wh.clientCfg)
	wh.hub.HandleClient(client)
}

//WSRoute returns provider that gives the web socket routes
func WSRoute(hub ws.Hub, cfg cfg.WSClient) *routeutils.Route {
	handler := wsHandler{
		hub:       hub,
		clientCfg: cfg,
	}

	return &routeutils.Route{
		Pattern: "/ws",
		Handler: &handler,
		Method:  http.MethodConnect,
	}
}
