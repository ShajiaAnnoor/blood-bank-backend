package queue

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type Queue interface {
	Push([]byte) error
	WaitPop() ([]byte, error)
	Close() error
}

type node struct {
	data []byte
	next *node
}

type queue struct {
	head *node
	tail *node
	len  int
}

func (q *queue) length() int {
	return q.len
}

func (q *queue) push(data []byte) {
	node := node{
		data: data,
		next: nil,
	}

	if q.head == nil {
		q.head = &node
		q.tail = q.head
	} else {
		q.tail.next = &node
		q.tail = q.tail.next
	}

	q.len++
}

func (q *queue) top() []byte {
	if q.head == nil {
		return nil
	}

	return q.head.data
}

func (q *queue) pop() {
	if q.head == nil {
		return
	}

	q.head = q.head.next
	q.len--
}

type syncQueue struct {
	q     queue
	mxLen int

	dataIn  chan []byte
	dataOut chan []byte

	closed chan struct{}
	once   sync.Once
}

func (q *syncQueue) isClosed() bool {
	select {
	case <-q.closed:
		return true
	default:
		return false
	}
}

func (q *syncQueue) Close() error {
	q.once.Do(func() {
		close(q.closed)
	})

	return nil
}

func (q *syncQueue) WaitPop() ([]byte, error) {
	if q.isClosed() {
		return nil, Closed
	}

	data, ok := <-q.dataOut
	if !ok {
		return nil, Closed
	}

	return data, nil
}

func (q *syncQueue) Push(data []byte) (err error) {
	if q.isClosed() {
		return Closed
	}

	defer func() {
		prob := recover()
		if prob != nil {
			err = Closed
		}
	}()

	q.dataIn <- data
	return
}

func (q *syncQueue) adjustLen() {
	for q.q.length() > q.mxLen {
		q.q.pop()
	}
}

func (q *syncQueue) sync() {
	defer func() {
		prob := recover()
		if prob != nil {
			logrus.Error(prob)
		}
	}()

	defer close(q.dataIn)
	defer close(q.dataOut)

	for {
		if q.isClosed() {
			return
		}

		if q.q.length() == 0 {
			select {
			case data, ok := <-q.dataIn:
				if !ok {
					return
				}

				q.q.push(data)
			case <-q.closed:
				return
			}
		} else {
			select {
			case <-q.closed:
				return
			case data, ok := <-q.dataIn:
				if !ok {
					return
				}
				q.q.push(data)
				q.adjustLen()
			case q.dataOut <- q.q.top():
				q.q.pop()
			}
		}
	}
}

func NewQueue(mxLen int) Queue {
	q := syncQueue{
		dataIn:  make(chan []byte),
		dataOut: make(chan []byte),
		closed:  make(chan struct{}),
		mxLen:   mxLen,
		once:    sync.Once{},
	}

	go q.sync()

	return &q
}
