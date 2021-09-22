package ws

import (
	"sync"
)

type ClientStore interface {
	Add(c Client)
	Remove(c Client)
	RangeByID(id string, f func(Client))
}

type clientStore struct {
	syncMap sync.Map
}

func (cs *clientStore) Add(c Client) {
	v, _ := cs.syncMap.LoadOrStore(c.ID(), &sync.Map{})
	clients := v.(*sync.Map)
	clients.Store(c, struct{}{})
}

func (cs *clientStore) Remove(c Client) {
	v, ok := cs.syncMap.Load(c.ID())
	if !ok {
		return
	}

	clients := v.(*sync.Map)
	clients.Delete(c)
}

func (cs *clientStore) RangeByID(id string, f func(Client)) {
	v, ok := cs.syncMap.Load(id)
	if !ok {
		return
	}

	clients := v.(*sync.Map)
	clients.Range(func(key interface{}, value interface{}) bool {
		c := key.(Client)
		f(c)
		return true
	})
}

func NewClientStore() ClientStore {
	return &clientStore{sync.Map{}}
}
