package pkg

import (
	"net"
	"sync"
	"sync/atomic"
)

type Backend struct {
	address net.Addr
	alive   bool
	mux     sync.RWMutex
}

func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.alive = alive
	b.mux.Unlock()
}
func (b *Backend) Address() net.Addr {
	return b.address
}

// IsAlive returns true when backend is alive
func (b *Backend) IsAlive() (alive bool) {
	b.mux.RLock()
	alive = b.alive
	b.mux.RUnlock()
	return
}

type ServerPool struct {
	backends []*Backend
	current  uint64
}

func (s *ServerPool) nextIndex() int {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
}
func (s *ServerPool) AddBackend(address net.Addr) {
	backend := &Backend{
		address: address,
		mux:     sync.RWMutex{},
	}
	backend.SetAlive(true)
	s.backends = append(s.backends, backend)
}
func (s *ServerPool) GetNextPeer() *Backend {
	next := s.nextIndex()
	l := len(s.backends) + next // start from next and move a full cycle
	for i := next; i < l; i++ {
		idx := i % len(s.backends) // take an index by modding with length
		if s.backends[idx].IsAlive() {
			if i != next {
				atomic.StoreUint64(&s.current, uint64(idx)) // mark the current one
			}
			return s.backends[idx]
		}
	}
	return nil
}
