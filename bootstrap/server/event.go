package server

import (
	"github.com/codeginga/locevt"
	"gitlab.com/Aubichol/hrishi-backend/container"
	"gitlab.com/Aubichol/hrishi-backend/event"
	"gitlab.com/Aubichol/hrishi-backend/ws"
)

//Event registers event related providers
func Event(c container.Container) {
	c.Register(func() locevt.Event {
		return locevt.NewEvent()
	})

	c.Resolve(func(e locevt.Event, hub ws.Hub) {
		e.Register(
			event.NameWSNotification,
			event.WSNotificationWorker(hub),
		)
	})
}
