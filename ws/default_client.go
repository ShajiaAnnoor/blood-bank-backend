package ws

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/hrishi-backend/cfg"
	"gitlab.com/Aubichol/hrishi-backend/pkg/queue"
	pkgsync "gitlab.com/Aubichol/hrishi-backend/pkg/sync"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type client struct {
	id string

	writeQ queue.Queue
	readQ  queue.Queue
	conn   *websocket.Conn

	mu        pkgsync.ReadWriteMutex
	kikedOnce sync.Once
	kiked     chan struct{}

	pingPeriod      time.Duration
	writeWait       time.Duration
	pongWait        time.Duration
	readMessageSize int64
}

func (c *client) Write(data []byte) (err error) {
	c.mu.Write(func() {
		if c.isKicked() {
			err = errors.New("kiked")
		}

		if err = c.writeQ.Push(data); err != nil {
			err = fmt.Errorf("write q push: %s", err.Error())
		}
	})

	return
}

func (c *client) Read() ([]byte, error) {
	data, err := c.readQ.WaitPop()
	if err != nil {
		return nil, fmt.Errorf("read q pop: %s", err.Error())
	}

	return data, nil
}

func (c *client) Kick() error {
	c.kikedOnce.Do(func() {
		close(c.kiked)
		c.writeQ.Close()
	})

	return nil
}

func (c *client) isKicked() bool {
	select {
	case <-c.kiked:
		return true
	default:
		return false
	}
}

func (c *client) ID() (id string) {
	c.mu.Read(func() {
		id = c.id
	})

	return
}

func (c *client) SetID(id string) {
	c.mu.Write(func() {
		c.id = id
	})
}

func (c *client) readPump() {
	defer func() {
		c.Kick()
		c.readQ.Close()
		c.conn.Close()
	}()

	c.conn.SetReadLimit(c.readMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(c.pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(c.pongWait)); return nil })

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Errorf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		if err = c.readQ.Push(message); err != nil {
			logrus.Error("readQ push: ", err)
			break
		}
	}
}

func (c *client) writePump() {
	ticker := time.NewTicker(c.pingPeriod)
	send := make(chan []byte)

	defer func() {
		c.Kick()
		ticker.Stop()
		c.conn.Close()
	}()

	go func(c chan<- []byte, q queue.Queue) {
		defer close(c)

		for {
			data, err := q.WaitPop()
			if err != nil {
				logrus.Error(err)
				return
			}

			c <- data
		}
	}(send, c.writeQ)

	for {
		select {
		case <-c.kiked:
			return
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(c.writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logrus.Error("ws write message: %s", err.Error())
				return
			}
		case data, ok := <-send:
			c.conn.SetWriteDeadline(time.Now().Add(c.writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				logrus.Error("next writer: ", err)
				return
			}

			_, err = w.Write(data)
			if err != nil {
				logrus.Error("writer write: ", err)
				return
			}

			if err = w.Close(); err != nil {
				logrus.Error("writer close: ", err)
				return
			}
		}
	}
}

func NewDefaultClient(conn *websocket.Conn, cfg cfg.WSClient) Client {
	c := client{
		kiked:           make(chan struct{}),
		conn:            conn,
		writeQ:          queue.NewQueue(cfg.WriteQBuffer),
		readQ:           queue.NewQueue(cfg.ReadQBuffer),
		pingPeriod:      cfg.PingPeriod,
		writeWait:       cfg.WriteWait,
		pongWait:        cfg.PongWait,
		readMessageSize: cfg.ReadMessageSize,
	}

	go c.readPump()
	go c.writePump()

	return &c
}
