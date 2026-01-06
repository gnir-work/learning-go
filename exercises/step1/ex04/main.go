/*
*
# Exercise 1.4: Struct Design & Methods (30 min)
Practice idiomatic struct design patterns.

## Task: Design a ConnectionPool struct for managing connections

Use pointer vs value receivers appropriately
Implement constructor pattern (NewConnectionPool())
Add methods: Get(), Put(), Close()
Use sync.Mutex for thread-safety
Demonstrate zero values working sensibly where possible
Focus: Constructor pattern, pointer vs value receivers, zero values, embedding
*/
package main

import (
	"fmt"

	"github.com/gnir-work/learning-go/exercises/step1/ex04/pool"
)

type PGConnection struct {
}

func (c *PGConnection) Close() error {
	fmt.Printf("Closing connection\n")
	return nil
}

func (c *PGConnection) ExecutePSQL(sql string) string {
	return fmt.Sprintf("Executed sql %q", sql)
}

func newPGConnection() *PGConnection {
	return &PGConnection{}
}

func main() {
	p := pool.NewConnectionPool(5, newPGConnection)
	defer func() {
		_ = p.Close()
	}()

	con, err := p.Get()
	if err != nil {
		fmt.Printf("Failed to fetch connection %v", err)
	}
	sqlResponse := con.ExecutePSQL("select * from tenant_123.person")
	fmt.Printf("Got response: %q\n", sqlResponse)
}
