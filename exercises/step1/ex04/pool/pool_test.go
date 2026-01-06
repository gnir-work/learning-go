package pool

import (
	"errors"
	"testing"
	"time"
)

type TestConnection struct {
	Closed bool
}

func (c *TestConnection) Close() error {
	c.Closed = true
	return nil
}

func testConnectionFactory() *TestConnection {
	return &TestConnection{
		Closed: false,
	}
}

func getConnectionFromPool(t *testing.T, p *ConnectionPool[*TestConnection]) *TestConnection {
	c, err := p.Get()
	if err != nil {
		t.Fatalf("Geto error %v when getting connection", err)
	}

	if c.Closed {
		t.Fatal("Connection was closed when received from the pool")
	}
	return c
}

func TestConnectionPool_HappyFlow(t *testing.T) {
	pool := NewConnectionPool(2, testConnectionFactory)
	getConnectionFromPool(t, pool)
	getConnectionFromPool(t, pool)
	err := pool.Close()
	if err != nil {
		t.Fatalf("Got error %v when closing the pool", err)
	}
}

func TestConnectionPool_ReturningConnectionToThePool(t *testing.T) {
	pool := NewConnectionPool(2, testConnectionFactory)
	c := getConnectionFromPool(t, pool)
	err := pool.Put(c)
	if err != nil {
		t.Fatalf("Error in returning connection to the pool %v", err)
	}
}

func TestConnectionPool_ReceivingTooManyConnectionsWithoutTimeout(t *testing.T) {
	pool := NewConnectionPool(2, testConnectionFactory)
	getConnectionFromPool(t, pool)
	c := getConnectionFromPool(t, pool)
	go func() {
		time.Sleep(2 * time.Second)
		pool.Put(c)
	}()
	getConnectionFromPool(t, pool)
}

func TestConnectionPool_ReceivingTooManyConnectionsTimesOut(t *testing.T) {
	pool := NewConnectionPool(2, testConnectionFactory, WithTimeout[*TestConnection](time.Second))
	getConnectionFromPool(t, pool)
	getConnectionFromPool(t, pool)
	_, err := pool.Get()
	if err == nil {
		t.Fatalf("Expected to receive error when getting third connection from the pool")
	}

	if !errors.Is(err, ErrPoolTimeout) {
		t.Fatalf("Expected error to be of type 'ErrPoolTimeout'")
	}
}
