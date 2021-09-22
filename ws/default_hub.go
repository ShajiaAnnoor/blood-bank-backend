package ws

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type hub struct {
	clientStore ClientStore
	handlers    map[string]Handler
}

func (h *hub) Send(userID string, data []byte) error {
	h.clientStore.RangeByID(userID, func(c Client) {
		if err := c.Write(data); err != nil {
			logrus.Error(fmt.Errorf("client write: %w", err))
		}
	})

	return nil
}

func (h *hub) kickClient(c Client) {
	if err := c.Kick(); err != nil {
		logrus.Error("client kick: ", err)
	}

	h.clientStore.Remove(c)
}

func (h *hub) checkAuth(c Client) {
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()

	<-timer.C
	if c.ID() == "" {
		h.kickClient(c)
	}
}

func (h *hub) handleRequest(c Client, request *RequestDTO) {
	handler, ok := h.handlers[request.Kind.String()]
	if !ok {
		c.Kick()
		return
	}

	handler.Handle(c, request)
}

func (h *hub) HandleClient(c Client) {
	go h.checkAuth(c)

	f := func(c Client) {
		defer h.kickClient(c)

		for {
			bytes, err := c.Read()
			if err != nil {
				logrus.Error("client read: ", err)
				return
			}

			request := RequestDTO{}
			if err := request.FromBytes(bytes); err != nil {
				logrus.Error("frombytes: ", err)
				return
			}

			h.handleRequest(c, &request)
		}
	}

	go f(c)
}

func NewHub(params struct {
	dig.In

	AuthHandler Handler `name:"auth"`
	ClientStore ClientStore
}) Hub {
	return &hub{
		clientStore: params.ClientStore,
		handlers: map[string]Handler{
			KindAuthentication.String(): params.AuthHandler,
		},
	}
}
