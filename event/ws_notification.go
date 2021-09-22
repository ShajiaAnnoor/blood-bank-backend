package event

import (
	"github.com/codeginga/locevt"
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/hrishi-backend/ws"
)

const (
	NameWSNotification string = "name-ws-notification"
)

type WSNotificationData struct {
	UserID string
	Data   []byte
}

func WSNotificationWorker(hub ws.Hub) func(locevt.Task) {
	return func(tsk locevt.Task) {
		data, ok := tsk.Data().(*WSNotificationData)
		if !ok {
			logrus.Error("could not convert to WSNotificationData")
			return
		}

		if err := hub.Send(data.UserID, data.Data); err != nil {
			logrus.Error("hub send: ", err)
			tsk.Retry()
		}
	}
}
