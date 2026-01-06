package pool

import (
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"
)

var ErrPoolTimeout = errors.New("connection pool timeout")

type Connection interface {
	Close() error
	comparable
}

type ConnectionPool[T Connection] struct {
	mutex             sync.Mutex
	connections       chan T
	timeout           time.Duration
	used              []T
	connectionFactory func() T
}

type Option[T Connection] func(pool *ConnectionPool[T])

func NewConnectionPool[T Connection](
	size uint32,
	connectionFactory func() T,
	opts ...Option[T],
) *ConnectionPool[T] {
	pool := &ConnectionPool[T]{
		mutex:             sync.Mutex{},
		connectionFactory: connectionFactory,
		timeout:           0,
		used:              make([]T, 0, size),
		connections:       make(chan T, size),
	}
	for _, opt := range opts {
		opt(pool)
	}

	for range size {
		pool.connections <- connectionFactory()
	}
	return pool
}

func WithTimeout[T Connection](timeout time.Duration) func(pool *ConnectionPool[T]) {
	return func(pool *ConnectionPool[T]) {
		pool.timeout = timeout
	}
}

func (p *ConnectionPool[T]) Get() (T, error) {

	var timeout <-chan time.Time = nil
	if p.timeout > 0 {
		timeout = time.After(p.timeout)
	}

	for {
		select {
		case con := <-p.connections:
			p.mutex.Lock()
			p.used = append(p.used, con)
			p.mutex.Unlock()
			return con, nil
		case <-timeout:
			var empty T
			return empty, fmt.Errorf("get connection timed out after %v: %w", p.timeout, ErrPoolTimeout)
		}
	}
}

func (p *ConnectionPool[T]) Put(conn T) error {
	p.connections <- conn

	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Find the first index of the value
	index := slices.Index(p.used, conn)
	if index == -1 {
		return fmt.Errorf("connection %v is not part of the used connections of the pool", conn)
	}

	p.used = slices.Delete(p.used, index, index+1)
	return nil
}

func (p *ConnectionPool[T]) Close() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, con := range p.used {
		p.connections <- con
	}
	p.used = make([]T, 0, len(p.used))

	close(p.connections)
	for con := range p.connections {
		if err := con.Close(); err != nil {
			return err
		}
	}
	return nil
}
